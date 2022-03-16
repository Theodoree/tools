package _00_199


/*
146. LRU缓存机制

运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

进阶:

你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例:

LRUCache cache = new LRUCache( 2  缓存容量  );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
*/

// 时间复杂度较高,空间复杂度较低
//type LRUCache struct {
//    Value       map[int]int
//    ValArr      []int
//    Cap         int
//    KeyIndexMap map[int]int
//    lru         map[int]int
//    lruCNT      int
//}
//
//func Constructor(capacity int) LRUCache {
//    return LRUCache{ValArr: make([]int, 0, capacity), Cap: capacity, KeyIndexMap: make(map[int]int), lru: make(map[int]int)}
//}
//
//func (this *LRUCache) Get(key int) int {
//    if idx, ok := this.KeyIndexMap[key]; ok {
//        delete(this.lru, key)
//        this.lru[key] = this.lruCNT
//        this.lruCNT++
//        return this.ValArr[idx]
//    }
//
//    return -1
//}
//
//func (this *LRUCache) Put(key int, value int) {
//
//    // 如果有则更新权重
//    if idx, ok := this.KeyIndexMap[key]; ok {
//        this.ValArr[idx] = value
//
//        // 如果存在开始淘汰
//    } else if len(this.ValArr) >= this.Cap {
//        var minCnt, deleteKey int
//        minCnt = this.lruCNT
//        for k, v := range this.lru {
//            if v < minCnt {
//                minCnt = v
//                deleteKey = k
//            }
//        }
//        delete(this.lru, deleteKey)
//        i := this.KeyIndexMap[deleteKey]
//        delete(this.KeyIndexMap, deleteKey)
//        this.KeyIndexMap[key] = i
//        this.ValArr[i] = value
//    } else {
//        this.ValArr = append(this.ValArr, value)
//        this.KeyIndexMap[key] = len(this.ValArr) - 1
//
//    }
//    this.lru[key] = this.lruCNT
//    this.lruCNT++
//
//}
//



// 使用node做 双向链表
type LRUCache struct {
    KeyMap   map[int]*Node
    Head     *Node
    Tail     *Node
    Capacity int
}

type Node struct {
    Key   int
    Value int
    Prev  *Node
    Next  *Node
}

func (this *LRUCache) MoveTail(cur *Node) {

    if this.Tail == cur {
        return
    }

    var (
        parent = cur.Prev
        next   = cur.Next
    )
    cur.Next = nil

    if parent != nil {
        parent.Next = next
        if next != nil {
            next.Prev = parent
        }
    }

    cur.Prev = this.Tail
    this.Tail.Next = cur
    this.Tail = cur
}

func (this *LRUCache) DeleteFirst() {

    deleteNode := this.Head.Next
    next := deleteNode.Next
    delete(this.KeyMap, deleteNode.Key)
    this.Head.Next = next
    if next != nil {
        next.Prev = this.Head
    }else{
        this.Tail = this.Head
    }


}
func Constructor(capacity int) LRUCache {
    n := &Node{}

    return LRUCache{KeyMap: make(map[int]*Node), Head: n, Tail: n, Capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
    if n, ok := this.KeyMap[key]; ok {
        this.MoveTail(n)
        return n.Value
    }

    return -1
}

func (this *LRUCache) Put(key int, value int) {

    if node, ok := this.KeyMap[key]; ok {
        node.Value = value
        this.MoveTail(node)
    } else {
        node := &Node{
            Key:   key,
            Value: value,
            Prev:  this.Tail,
            Next:  nil,
        }

        if len(this.KeyMap) >= this.Capacity {
            this.DeleteFirst()
        }
        this.Tail.Next = node
        node.Prev = this.Tail
        this.Tail = node

        this.KeyMap[key] = node
    }
}

