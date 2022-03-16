package worker

import (
    "errors"
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

const (
    defaultMaxGoroutine   = 50
    defaultMinPushChan    = 100
    defaultMaxPushChan    = 10000
    defaultWorkerChanRate = 0.1

    prepareClose = 1
    closed       = 2
)

type Pool struct {
    pushChan chan *task
    exitChan chan struct{}
    taskList []*task // 任务

    taskPool *taskPool // sync.pool

    worker        []*Worker
    workerChanCap int
    lastIndex     int

    freeList *chainList

    close uint64
}

var (
    isCloseError = errors.New("slow close.. don't send any task ")
)

func (p *Pool) insertFreeWorker(worker *Worker) {
    p.freeList.pushWorker(worker)
}

func (p *Pool) PushTask(fn func(), panicHandler func(err error)) error {
    if atomic.LoadUint64(&p.close) == prepareClose || atomic.LoadUint64(&p.close) == closed {
        return isCloseError
    }

    task := p.taskPool.get()
    task.fn = fn
    task.panicHandler = panicHandler
    if atomic.LoadUint64(&p.close) == prepareClose || atomic.LoadUint64(&p.close) == closed {
        return isCloseError
    }
    p.pushChan <- task
    return nil
}

func (p *Pool) scheduling() {

    ticker := time.NewTicker(time.Second * 2)

    for {
        select {
        case t := <-p.pushChan:
            // 如果有空闲的worker直接运行
            worker := p.freeList.popWorker()
            // 不为空直接消费
            if worker != nil {
                worker.taskChan <- t
                break
            }

            curIndex := p.lastIndex
            isScan := false

            for {
                if len(p.worker[curIndex].taskChan) < cap(p.worker[curIndex].taskChan) {
                    p.worker[curIndex].taskChan <- t
                    p.lastIndex = curIndex + 1
                    break
                }

                curIndex++
                // 如果已经扫描一遍了,没有空闲的worker,那么先放入数组
                if isScan && curIndex == p.lastIndex {
                    p.taskList = append(p.taskList, t)
                    break
                }

                // 重置下标
                if curIndex == len(p.worker) {
                    curIndex = 0
                    isScan = true
                }

            }
        case <-ticker.C:
            curIndex := p.lastIndex
            scanTimes := 0
            for len(p.taskList) > 0 {
                if len(p.worker[curIndex].taskChan) < cap(p.worker[curIndex].taskChan) {
                    p.worker[curIndex].taskChan <- p.taskList[len(p.taskList)-1]
                    p.taskList = p.taskList[:len(p.taskList)-1]
                }
                curIndex++

                if scanTimes == p.workerChanCap && curIndex == p.lastIndex {
                    break
                }

                if curIndex == len(p.worker) {
                    curIndex = 0
                    scanTimes++
                }
            }


        case <-p.exitChan:

            // 关闭所有worker
            for idx := range p.worker {
                p.worker[idx].close()
            }

            var closeNum uint64

            for len(p.worker) > 0 {
                if atomic.LoadUint64(&p.worker[0].step) == closed {
                    closeNum++
                    p.worker = p.worker[1:]
                }
                // 睡1ms
                doSleep()
            }

            close(p.exitChan)
            close(p.pushChan)
            ticker.Stop()
            p.worker = nil
            p.freeList = nil
            p.doResetTask()
            p.taskList = nil

            return
        }

        // 防止越界
        if p.lastIndex == len(p.worker) {
            p.lastIndex = 0
        }
    }

}

func (p *Pool) Close() {
    // 阻塞chan,等待所有任务执行完毕
    atomic.StoreUint64(&p.close, prepareClose)
    p.exitChan <- struct{}{}
}

func (p *Pool) doResetTask() {

    var c = make(chan struct{}, 20)
    wait := sync.WaitGroup{}

    // 多线程执行剩余的任务
    for _, v := range p.taskList {
        c <- struct{}{}
        wait.Add(1)
        go func(task *task) {
            defer func() {
                if err := recover(); err != nil && task.panicHandler != nil {
                    buf := stack(3)
                    task.panicHandler(errors.New(fmt.Sprintf("%s", buf)))
                }
                <-c
                wait.Done()
            }()
            // 执行方法
            task.fn()
        }(v)
    }
    wait.Wait()

}

func NewWorkPool(workerNum int, pushChanCap int) (*Pool, error) {
    // 一些参数检查,限制携程数、
    if workerNum > defaultMaxGoroutine {
        workerNum = defaultMaxGoroutine // 最多一百个携程
    }

    // pushChan cap
    if pushChanCap > defaultMaxPushChan {
        pushChanCap = defaultMaxPushChan
    }

    if pushChanCap < defaultMinPushChan {
        pushChanCap = defaultMinPushChan
    }
    if pushChanCap%10 != 0 {
        return nil, errors.New("pushChanCap大小必须为十的倍数")
    }

    pool := &Pool{
        pushChan:      make(chan *task, pushChanCap),
        exitChan:      make(chan struct{}),
        workerChanCap: int(float64(pushChanCap) * defaultWorkerChanRate),
        taskPool:      newTaskPool(),
        worker:        make([]*Worker, 0, workerNum),
    }
    pool.freeList = newChainList()

    for i := 0; i < workerNum; i++ {
        worker := newWorker(pool.workerChanCap)
        worker.pool = pool
        pool.freeList.pushWorker(worker)
        pool.worker = append(pool.worker, worker)
        go worker.Run()

    }

    go pool.scheduling()

    return pool, nil
}
