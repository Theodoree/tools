// Copyright 2012 Google Inc. All rights reserved.
// Author: Ric Szopa (Ryszard) <ryszard.szopa@gmail.com>

// Package skiplist implements skip list based maps and sets.
//
// Skip lists are a data structure that can be used in place of
// balanced trees. Skip lists use probabilistic balancing rather than
// strictly enforced balancing and as a result the algorithms for
// insertion and deletion in skip lists are much simpler and
// significantly faster than equivalent algorithms for balanced trees.
//
// Skip lists were first described in Pugh, William (June 1990). "Skip
// lists: a probabilistic alternative to balanced
// trees". Communications of the ACM 33 (6): 668–676

package skiplist

import "math/rand"

const (
    p        = 0.25
    maxLevel = 32
)

/*
tip:            A  B
    条件A: 如果key小于节点B,插入位置必定从节点A开始
*/

type SkipList struct {
    lessThan func(l, r interface{}) bool // 比较方法   l < r  true
    header   *node
    footer   *node
    length   int
    Level    int
}

func (s *SkipList) Len() int {
    return s.length
}

// 根据条件A寻找指定节点
func (s *SkipList) getPath(current *node, update []*node, key interface{}) *node {
    depth := len(current.forward) - 1 // 从顶部开始扫描,一直扫描到bottom

    for i := depth; i >= 0; i-- {
        for current.forward[i] != nil && s.lessThan(current.forward[i].key, key) { // 寻找插入位置,直到forward[i]不满足条件A,只到当前节点大于key
            current = current.forward[i]
        }
        if update != nil {
            update[i] = current
        }
    }
    return current.next()
}

// 构造器
func (s *SkipList) Iterator() Iterator {
    return &iter{
        current: s.header,
        list:    s,
    }
}

// seek返回一个迭代器
func (s *SkipList) Seek(key interface{}) Iterator {
    current := s.getPath(s.header, nil, key) // 从头部遍历
    if current == nil {
        return nil
    }
    return &iter{
        current: current,
        key:     current.key,
        list:    s,
        value:   current.value,
    }
}

// 从头部迭代
func (s *SkipList) SeekToFirst() Iterator {
    if s.length == 0 {
        return nil
    }

    current := s.header.next()

    return &iter{
        current: current,
        key:     current.key,
        list:    s,
        value:   current.value,
    }
}

// 从尾部迭代
func (s *SkipList) SeekToLast() Iterator {
    current := s.footer
    if current == nil {
        return nil
    }

    return &iter{
        current: current,
        key:     current.key,
        list:    s,
        value:   current.value,
    }
}

// range迭代器
func (s *SkipList) Range(from, to interface{}) Iterator {
    start := s.getPath(s.header, nil, from)

    return &rangeIterator{
        iter: iter{
            current: &node{
                forward:  []*node{start},
                backward: start,
            },
            list: s,
        },
        upperLimit: to,
        lowerLimit: from,
    }
}

func (s *SkipList) level() int {
    return len(s.header.forward) - 1
}

func (s *SkipList) effectiveMaxLevel() int {
    return maxInt(s.level(), s.Level)
}

// 返回一个新的random等级
func (s SkipList) randomLevel() (n int) {
    for n = 0; n < s.effectiveMaxLevel() && rand.Float64() < p; n++ {
    }
    return
}

// 获取value
func (s *SkipList) Get(key interface{}) (value interface{}, ok bool) {
    candidate := s.getPath(s.header, nil, key)

    if candidate == nil || candidate.key != key {
        return nil, false
    }

    return candidate.value, true
}

// 获取更大符合条件的节点
func (s *SkipList) GetGreaterOrEqual(min interface{}) (actualKey, value interface{}, ok bool) {
    candidate := s.getPath(s.header, nil, min)

    if candidate != nil {
        return candidate.key, candidate.value, true
    }
    return nil, nil, false
}

// set函数
func (s *SkipList) Set(key, value interface{}) {
    if key == nil {
        panic("goskiplist: nil keys are not supported")
    }

    // level从零开始,所以需要分配一个长度
    update := make([]*node, s.level()+1, s.effectiveMaxLevel()+1)
    candidate := s.getPath(s.header, update, key) // 寻找插入位置

    if candidate != nil && candidate.key == key {
        candidate.value = value
        return
    }

    // 获取一个随机level
    newLevel := s.randomLevel()

    if currentLevel := s.level(); newLevel > currentLevel {

        // 如果随机level大于当前level,那么填充数组
        for i := currentLevel + 1; i <= newLevel; i++ {
            update = append(update, s.header)
            s.header.forward = append(s.header.forward, nil)
        }
    }

    newNode := &node{
        forward: make([]*node, newLevel+1, s.effectiveMaxLevel()+1),
        key:     key,
        value:   value,
    }

    // 将原来的插入位置,设置为当前节点的后继
    if previous := update[0]; previous.key != nil {
        newNode.backward = previous
    }

    // 更新途中略过的node,并且父子节点相互更新
    for i := 0; i <= newLevel; i++ {
        newNode.forward[i] = update[i].forward[i]
        update[i].forward[i] = newNode
    }

    s.length++

    // 将父节点的后继设置为自己(上面已经将旧尾部设置为当前节点的后继)
    if newNode.forward[0] != nil {
        if newNode.forward[0].backward != newNode {
            newNode.forward[0].backward = newNode
        }
    }

    // 检查是否为新尾部
    if s.footer == nil || s.lessThan(s.footer.key, key) {
        s.footer = newNode
    }
}

func (s *SkipList) Delete(key interface{}) (value interface{}, ok bool) {
    if key == nil {
        panic("goskiplist: nil keys are not supported")
    }
    update := make([]*node, s.level()+1, s.effectiveMaxLevel())
    candidate := s.getPath(s.header, update, key)

    if candidate == nil || candidate.key != key {
        return nil, false
    }

    previous := candidate.backward
    // 如果footer等于待删除节点,那么更新footer
    if s.footer == candidate {
        s.footer = previous
    }

    next := candidate.next()
    // 更新前驱节点的后继成员
    if next != nil {
        next.backward = previous
    }

    // 更新途中节点
    for i := 0; i <= s.level() && update[i].forward[i] == candidate; i++ {
        update[i].forward[i] = candidate.forward[i]
    }

    // 检查下面几层是否为空
    for s.level() > 0 && s.header.forward[s.level()] == nil {
        s.header.forward = s.header.forward[:s.level()]
    }
    s.length--

    return candidate.value, true
}

func NewCustomMap(lessThan func(l, r interface{}) bool) *SkipList {
    return &SkipList{
        lessThan: lessThan,
        header: &node{
            forward: []*node{nil},
        },
        Level: maxLevel,
    }
}



type Ordered interface {
    LessThan(other Ordered) bool
}

func New() *SkipList {
    comparator := func(left, right interface{}) bool {
        return left.(Ordered).LessThan(right.(Ordered))
    }
    return NewCustomMap(comparator)

}
func NewIntMap() *SkipList {
    return NewCustomMap(func(l, r interface{}) bool {
        return l.(int) < r.(int)
    })
}


func NewStringMap() *SkipList {
    return NewCustomMap(func(l, r interface{}) bool {
        return l.(string) < r.(string)
    })
}
