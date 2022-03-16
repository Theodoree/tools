package _go

import "sync"

/*
面试题 16.25. LRU缓存
设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。
缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。
当缓存被填满时，它应该删除最近最少使用的项目。

它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。


*/

type LRUCache struct {
    KeyMap   map[int]*Node
    Head     *Node
    Tail     *Node
    Capacity int
    pool     *sync.Pool
}

type Node struct {
    Key   int
    Value int
    Prev  *Node
    Next  *Node
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        KeyMap: make(map[int]*Node),
    }
    l.pool = &sync.Pool{}
    l.pool.New = func() interface{} {
        return &Node{}
    }
    l.Capacity = capacity
    return l
}

func (cache *LRUCache) Get(key int) int {
    n, ok := cache.KeyMap[key]
    if ok {
        cache.move2Head(n)
        return n.Value
    }
    return - 1

}

func (cache *LRUCache) push2Head(n *Node) {
    /*
       case:填充尾部和head
    */

    if !cache.check(n) {
        return
    }


    oldHead := cache.Head
    n.Next = oldHead
    oldHead.Prev = n

    cache.Head = n

}
func (cache *LRUCache) move2Head(n *Node) {
    /*
       case:如果已经在头部 直接返回
    */
    if n == cache.Head {
        return
    }

    /*
       case:填充尾部和head
    */

    if !cache.check(n) {
        return
    }

    /*
       case:先处理一遍
    */

    parent := n.Prev
    next := n.Next
    if parent != nil {
        parent.Next = next
    }

    if next != nil {
        next.Prev = parent
    }

    if n == cache.Tail {
        cache.Tail = parent
    }

    n.Prev = nil

    oldHead := cache.Head
    n.Next = oldHead
    oldHead.Prev = n

    cache.Head = n

}


func (cache *LRUCache) removeLast() {

    // 删除旧尾指针
    n := cache.Tail
    if n != nil {
        delete(cache.KeyMap, n.Key)
        old:=cache.Tail
        parent := cache.Tail.Prev
        if parent != nil {
            parent.Next = nil
        }
        cache.Tail = parent

        cache.pool.Put(old)
    }
}

func (cache *LRUCache) check(n *Node) bool {
    if cache.Head == nil {
        cache.Head = n
        cache.Tail = n
        return false
    }

    return true
}

func (cache *LRUCache) Put(key int, value int) {
    node, ok := cache.KeyMap[key]
    if !ok {

        // 删除最少使用的node
        if cache.Capacity <= len(cache.KeyMap) {
            cache.removeLast()
        }

        // move到尾部
        n := cache.pool.Get().(*Node)
        n.Key = key
        n.Value = value
        cache.KeyMap[key] = n
        cache.push2Head(n)

    }else{
        node.Value = value
        cache.move2Head(node)
    }

}

func (cache *LRUCache) print() {
    c := cache.Head
    for c != nil {
        print(c.Key, " ")
        c = c.Next
    }
    println()
}

