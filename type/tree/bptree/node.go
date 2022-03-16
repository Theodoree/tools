package bptree

import "unsafe"

type node struct {
    Parent    *node            // 8byte
    child     []unsafe.Pointer // 防止非叶子节点分配大量内存  24byte  如果是叶子节点,那么child就是value
    childKeys []uint64         // 这里保存子节点的key值  24byte
    numKeys   int
    IsLeaf    bool // 是否为叶子节点            1byte
}


func newNode() *node {
    var n = &node{}

    n.childKeys = make([]uint64, order - 1, order - 1)
    n.child = make([]unsafe.Pointer, order, order)
    n.IsLeaf = false
    n.Parent = nil
    return n

}

func newLeafNode() *node {
    n := newNode()
    n.IsLeaf = true
    return n
}

func (n *node) reset() {
    for idx := range n.childKeys {
        n.childKeys[idx] = 0
    }
    for idx := range n.child {
        n.child[idx] = nil
    }
    n.Parent = nil
    n.numKeys = 0
}
