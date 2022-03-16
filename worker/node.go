package worker

import (
    "sync"
    "sync/atomic"
)

type chainList struct {
    mutex  sync.Mutex
    header *node
    tail   *node
    pool   *nodePool
    atomic int64
}

func (n *chainList) pushWorker(worker *Worker) {
    node := n.pool.get()
    n.mutex.Lock()
    if n.tail != nil {
        node.val = worker
        n.tail.next = node
        n.tail = node
    } else {
        n.tail = node
        n.header = node
    }
    atomic.AddInt64(&n.atomic, 1)
    n.mutex.Unlock()
}

func (n *chainList) popWorker() *Worker {
    // fast path
    if !n.check() {
        return nil
    }

    n.mutex.Lock()
    // 第二次检查
    if !n.check() {
        return nil
    }
    headerNode := n.header
    if headerNode.next != nil {
        n.header = headerNode.next
    } else {
        n.header = nil
        n.tail = nil
    }

    worker := headerNode.val
    n.pool.put(headerNode)

    atomic.AddInt64(&n.atomic, -1)
    n.mutex.Unlock()

    return worker
}

func (n *chainList) check() bool {
    return atomic.LoadInt64(&n.atomic) > 0
}

func newChainList() *chainList {
    chainList := &chainList{}
    chainList.mutex = sync.Mutex{}
    chainList.pool = newNodePool()
    return chainList
}

type node struct {
    val  *Worker
    next *node
}

func (n *node) reset() {
    n.val = nil
    n.next = nil
}

type nodePool struct {
    pool sync.Pool
}

func (p *nodePool) get() *node {
    return p.pool.Get().(*node)
}

func (p *nodePool) put(n *node) {
    n.reset()
    p.pool.Put(n)
}

func newNodePool() *nodePool {
    p := &nodePool{}
    p.pool.New = func() interface{} {
        return &node{}
    }
    return p
}
