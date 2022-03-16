package decorator

import (
    "context"
    "sync"
)

/*
简单的线程池
 */

type ThreadPool interface {
    InitThreadNumber(num int)
    SendNewFunc(fc func())
}

type threadPool struct {
    sync.Mutex
    Work     []*work
    CurIndex int
    WorkLen  int
}

func NewThreadPool(num int) *threadPool {
    pool := &threadPool{WorkLen: num}
    pool.InitThreadNumber(num)

    return pool
}

func (t *threadPool) InitThreadNumber(num int) {

    t.Work = make([]*work, 0, num)

    for i := 0; i < num; i++ {
        t.Work = append(t.Work, &work{
            WorkChan: make(chan func(), 100),
            Context:  context.Background(),
        })
    }

    // 开启线程
    for _, v := range t.Work {
        go v.Start()
    }

}

func (t *threadPool) SendNewFunc(fc func()) {
    t.Lock()
    for {
        if len(t.Work[t.CurIndex].WorkChan) < cap(t.Work[t.CurIndex].WorkChan) {
            // 这里不会阻塞
            t.Work[t.CurIndex].WorkChan <- fc
            t.curIndexInc()
            t.Unlock()
            return
        } else {
            // 这里会
            c := t.Work[t.CurIndex].WorkChan
            t.curIndexInc()
            t.Unlock()
            c <- fc
            return
        }
    }

}

func (t *threadPool) curIndexInc() {
    t.CurIndex++
    if t.CurIndex == t.WorkLen {
        t.CurIndex = 0
    }
}

type work struct {
    WorkChan chan func()
    Context  context.Context
    sync.Mutex
}

func (w *work) Start() {
    for {
        select {
        case job := <-w.WorkChan:
            w.Do(job)
        case <-w.Context.Done():
            return

        }
    }
}
func (w *work) Do(job func()) {
    defer func() {
        if err := recover(); err != nil {
            println(err)
        }
    }()
    job()
}
func (w *work) Done() {
    //blblbl....
}


type ThreadPoolWithNonblocking interface {
    InitThreadNumber(num int)
    SendNewFunc(fc func())
}

// .... 然后下面一堆实现,重写SendNewFunc方法 blblbl 大概就是怎么个东西。