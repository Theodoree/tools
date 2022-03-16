package btree

import "sync"

// syncpool
type FreeList struct {
    pool sync.Pool
}

func NewFreeList() *FreeList {

    f := &FreeList{}

    f.pool.New = func() interface{} {
        return new(node)
    }

    return f
}

func (f *FreeList) newNode() (n *node) {
    return f.pool.New().(*node)
}

func (f *FreeList) freeNode(n *node) () {
    f.pool.Put(n)
}
