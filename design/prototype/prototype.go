package prototype

import "sync"

// Prototype Design Pattern 原型模式
// 大致分为两种,一种是浅拷贝也就是引用类型只复制地址,第二种是深拷贝,递归拷贝,重新申请内存

type service struct {
    m *concurrentMap
}

func (s *service) clone() *service {
    return &service{
        m: s.m,
    }
}

func (s *service) cloneDeeply() *service {
    return &service{
        m: s.m.clone(),
    }
}

type concurrentMap struct {
    m     map[uint64]uint64
    mutex sync.Mutex
}

func NewConcurrentMap(m map[uint64]uint64) *concurrentMap {
    if m == nil {
        m = make(map[uint64]uint64)
    }

    return &concurrentMap{
        m:     m,
        mutex: sync.Mutex{}}
}

func (m *concurrentMap) clone() *concurrentMap {
    newMap := make(map[uint64]uint64)
    m.mutex.Lock()
    for k, v := range newMap {
        newMap[k] = v
    }
    m.mutex.Unlock()

    return NewConcurrentMap(newMap)
}

//... get delete set
