package rbtree

import "sync"

var _resourcesPool *resourcesPool

func init() {
    _resourcesPool = newResourcesPool()
}

type resourcesPool struct {
    nodePool sync.Pool
}

func newResourcesPool() *resourcesPool {
    pool := &resourcesPool{}
    pool.nodePool.New = func() interface{} {
        return newNode()
    }
    return pool
}

func (pool *resourcesPool) getNode() *node {
    return pool.nodePool.Get().(*node)
}

func (pool *resourcesPool) putNode(n *node) {
    n.reset()
    pool.nodePool.Put(n)
}
