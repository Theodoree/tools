package command

import (
    "sync"
    "time"
)

type Service interface {
    Connection(conn string) Client
    mainLoop()
    Start()
}
type service struct {
    clients []Client
    sync.Mutex
}

func NewService() Service {
    return &service{Mutex: sync.Mutex{}}
}
func (s *service) mainLoop() {
    var events = make([]*event, 0, 100)
    for {
        for _, v := range s.clients {
            // do something,从队列中读取请求 blblblbl
            event := v.GetEvents()
            if event != nil {
                events = append(events, event)
            }
        }

        var queue []command
        for _, event := range events {
            queue = append(queue, event.GetCommand())
        }

        for _, v := range queue {
            v.Invoke()
        }

        events = events[:0]
        //设置一个睡眠时间
        time.Sleep(time.Millisecond * 100)
    }

}
func (s *service) Connection(conn string) Client {
    client := NewClient(conn)
    s.Lock()
    s.clients = append(s.clients, client)
    s.Unlock()
    return client
}
func (s *service) Start() {
    go s.mainLoop()
}

type Client interface {
    GetEvents() *event
    PushEvents(*event)
}
type client struct {
    ip     string
    events []*event
    sync.Mutex
}

// chan嘛,可用可不用,里面也是锁,本身我也没有线程传递的需求。
func NewClient(ip string) Client {
    return &client{ip: ip, Mutex: sync.Mutex{}}
}
func (s *client) GetEvents() (val *event) {
    s.Lock()
    if len(s.events) > 0 {
        val = s.events[0]
        s.events = s.events[1:]
    }
    s.Unlock()
    return
}
func (s *client) PushEvents(event *event) {
    s.Lock()
    s.events = append(s.events, event)
    s.Unlock()
}

type command interface {
    Invoke()
}
type action uint64

const (
    printCmd = iota + 1
)

type printCommand struct {
    msg string
}

func (cmd *printCommand) Invoke() {
    println(cmd.msg)
}
func NewPrintCommand(msg string) command {
    return &printCommand{
        msg: msg,
    }
}

type event struct {
    Action action
    Data   string
}

func (e *event) GetCommand() command {
    switch e.Action {
    case printCmd:
        return NewPrintCommand(e.Data)
    }

    // 然后这里做处理
    return nil
}
