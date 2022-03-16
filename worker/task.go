package worker

import "sync"

type taskPool struct {
    pool sync.Pool
}

func (t *taskPool) get() *task {
    return t.pool.Get().(*task)
}

func (t *taskPool) put(task *task) {
    task.reset()
    t.pool.Put(task)
}

type task struct {
    fn           func()
    panicHandler func(err error)
}

func (t *task) reset() {
    t.fn = nil
    t.panicHandler = nil
}

func newTaskPool() *taskPool {
    pool := taskPool{}
    pool.pool.New = func() interface{} {
        return new(task)
    }

    return &pool
}
