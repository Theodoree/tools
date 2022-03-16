package rbtree

const (
    black = iota
    red
)

type node struct {
    parent *node
    child  [2]*node
    key    uint64
    value  uint64
    color  uint8
}

func newNode() *node {
    return &node{
        color: red,
    }
}

/*
desc:重置node
*/
func (n *node) reset() {
    for idx := range n.child {
        n.child[idx] = nil
    }
    n.key = 0
    n.value = 0
    n.color = red
}

/*
desc:key
*/
func (n *node) Key() uint64 {
    return n.key
}

/*
desc:返回value
*/
func (n *node) Value() uint64 {
    return n.value
}

/*
desc:返回父节点
*/
func (n *node) Parent() *node {
    return n.parent
}

/*
desc:返回叔节点
*/
func (n *node) Uncle() *node {
    grandParent := n.GrandParent()
    if grandParent == nil {
        return nil
    }

    if grandParent.child[leftNodeIndex] == n {
        return grandParent.child[rightNodeIndex]
    }
    return grandParent.child[leftNodeIndex]
}

/*
desc:返回兄弟节点
*/
func (n *node) sibling() *node {
    if n.parent == nil {
        return nil
    }

    if n.parent.child[leftNodeIndex] == n {
        return n.parent.child[rightNodeIndex]
    }
    return n.parent.child[leftNodeIndex]
}

/*
desc:返回祖先节点
*/
func (n *node) GrandParent() *node {
    if n.parent != nil {
        return n.parent.parent
    }
    return nil
}
