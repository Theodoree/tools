package worker

import (
    "errors"
    "fmt"
)

type Worker struct {
    pool     *Pool
    taskChan chan *task
    exitChan chan struct{}
    curTask  *task
    step     uint64 // 1: 发送消息 2:chan中的消息已被消费完毕
}

func (w *Worker) Run() {
    for {
        select {
        case task := <-w.taskChan:
            // 清理剩余资源
            w.curTask = task
            w.doWorker()
            if len(w.taskChan) == 0 { // 如果任务全部完成,则进入freeList
                w.pool.insertFreeWorker(w)
            }
        case <-w.exitChan:
            // 读取所有任务
            for len(w.taskChan) > 0 {
                task := <-w.taskChan
                w.curTask = task
                w.doWorker()
            }

            w.step = closed
            w.pool = nil
            w.curTask = nil
            close(w.exitChan)
            close(w.taskChan)
            return
        }
    }

}

func (w *Worker) doWorker() {

    defer func() {
        if err := recover(); err != nil && w.curTask.panicHandler != nil {
            buf := stack(3)
            w.curTask.panicHandler(errors.New(fmt.Sprintf("%s", buf)))
        }

        w.pool.taskPool.put(w.curTask)

    }()
    // 执行方法
    w.curTask.fn()
}

func (w *Worker) close() {
    w.step = prepareClose
    w.exitChan <- struct{}{}
}

func newWorker(cap int) *Worker {
    worker := &Worker{}
    worker.taskChan = make(chan *task, cap)
    worker.exitChan = make(chan struct{})
    return worker
}
