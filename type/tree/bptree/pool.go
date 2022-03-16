package bptree

import "sync"

var (
    defaultResourcePool = newResourcePool()
)

// 简单的syncpool
type resourcePool struct {
    nodePool  sync.Pool
    leafPool  sync.Pool
    valuePool sync.Pool
    tempKeyPool  sync.Pool
    tempPool  sync.Pool
}

func newResourcePool() *resourcePool {
    p := resourcePool{}
    p.nodePool.New = func() interface{} {
        return newNode()
    }

    p.leafPool.New = func() interface{} {
        return newLeafNode()
    }

    p.valuePool.New = func() interface{} {
        return NewValue()
    }

    p.tempPool.New = func() interface{} {
        return make([]interface{}, order+1)
    }

    return &p
}

func (n *resourcePool) GetNode() *node {
    node := n.nodePool.Get().(*node)
    return node
}

func (n *resourcePool) PutNode(node *node) {
    node.reset()
    n.nodePool.Put(node)
}

func (n *resourcePool) GetLeaf() *node {
    node := n.leafPool.Get().(*node)
    return node
}

func (n *resourcePool) PutLeaf(node *node) {
    node.reset()
    n.leafPool.Put(node)
}

func (n *resourcePool) GetValue() *value {
    val := n.valuePool.Get().(*value)
    return val
}

func (n *resourcePool) PutValue(val *value) {
    val.reset()
    n.valuePool.Put(val)
}

func (n *resourcePool) GetTempSlice() []interface{} {
    m := n.tempPool.Get().([]interface{})
    return m
}
func (n *resourcePool) PutTempSlice(s []interface{}) {
    for idx := range s {
        if s[idx] == nil {
            break
        }
        s[idx] = nil
    }
    n.tempPool.Put(s)
}
