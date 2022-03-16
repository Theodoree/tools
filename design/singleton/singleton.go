package singleton

import "sync"



// 懒汉式
var srv = &service{}

type service struct {
    id uint64
    sync.Once
}

func GetService() *service {
    srv.Do(func() {
        // do something
    })
    return srv
}


// 饿汉式
var srv1 = &service1{}

func init() {
    srv1.init()
}

type service1 struct {}

func GetService1() *service1 {
    return srv1
}

func (s *service1) init() {}
