package _00_499

import (
    "container/heap"
    "time"
)

/*
460. LFU缓存

设计并实现最不经常使用（LFU）缓存的数据结构。它应该支持以下操作：get 和 put。

get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
put(key, value) - 如果键不存在，请设置或插入值。当缓存达到其容量时，它应该在插入新项目之前，使最不经常使用的项目无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，最近最少使用的键将被去除。

进阶：
你是否可以在 O(1) 时间复杂度内执行两项操作？

示例：

LFUCache cache = new LFUCache( 2 /* capacity (缓存容量)  );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回 1
cache.put(3, 3);    // 去除 key 2
cache.get(2);       // 返回 -1 (未找到key 2)
cache.get(3);       // 返回 3
cache.put(4, 4);    // 去除 key 1
cache.get(1);       // 返回 -1 (未找到 key 1)
cache.get(3);       // 返回 3
cache.get(4);       // 返回 4
*/

type LFUCache struct {
    m   map[int]*item
    pq  PQ // 优先队列
    cap int
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        m:   make(map[int]*item, capacity),
        pq:  make(PQ, 0, capacity),
        cap: capacity,
    }
}

// Get 获取 key 的值
func (c *LFUCache) Get(key int) int {

    if it, ok := c.m[key]; ok {
        c.pq.update(it)
        return it.value
    }
    return -1
}

// Put 把 key， value 成对地放入 LFUCache
// 如果 LFUCache 已满的话，会删除掉 LFUCache 中使用最少的 key
func (c *LFUCache) Put(key int, value int) {
    if c.cap <= 0 {
        return
    }
    it, ok := c.m[key]
    // key 已存在，就更新 value
    if ok {
        c.m[key].value = value
        c.pq.update(it)
        return
    }

    it = &item{key: key, value: value}
    // pq 已满的话，需要先删除，再插入
    if len(c.pq) == c.cap {
        temp := heap.Pop(&c.pq).(*item)
        delete(c.m, temp.key)
    }
    c.m[key] = it
    heap.Push(&c.pq, it)
}

type item struct {
    key, value int

    frequence, index int
    date             time.Time
}

// 使用heap实现优先队列
type PQ []*item

func (pq PQ) Len() int {
    return len(pq)
}

// 淘汰使用次数最少 相当使用次数时 淘汰使用时间最早的
func (pq PQ) Less(i, j int) bool {
    if pq[i].frequence == pq[j].frequence {
        return pq[i].date.Before(pq[j].date)
    }
    return pq[i].frequence < pq[j].frequence
}

func (pq PQ) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
    n := len(*pq)
    item := x.(*item)
    item.index = n
    item.date = time.Now()
    *pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

// 更新date frequence
func (pq *PQ) update(item *item) {
    item.frequence++
    item.date = time.Now()
    heap.Fix(pq, item.index)
}
