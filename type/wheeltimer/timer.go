package wheeltimer

import (
    "github.com/Theodoree/sample_project/worker"
    "sync"
    "time"
)

// 时间轮

/*
因为是编译性语言,所以这里尽量不使用Interface,没必要增加运行时的开支
*/

type Srv struct {
    mutex  sync.Mutex
    taskId uint64 //当前已使用的id

    root     *dayWheel
    workPool worker.Pool
}

type dayWheel struct {
    weekDay    time.Weekday
    hour       uint8
    sec        uint8
    hourWheels [24]*hourWheel
}

type hourWheel struct {
    minuteWheels [60]*minuteWheel
}

type minuteWheel struct {
    work [60][]Task
}

type Task struct {
    Id        uint64 // taskId,可以用该id来删除定时任务
    fn        func() //  具体执行任务
    weekDay   int8   // -1 代表每日都执行
    threshold int    // 最多重试次数 -1代表重试到成功为止
    err       func() // 执行失败调用该函数
    suc       func() // 执行成功后调用该函数
}

func (s *Srv) Start() {
    go s.root.start()
}

func (w *dayWheel) start() {
}

