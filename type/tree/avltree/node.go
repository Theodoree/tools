package avltree

import "fmt"

type node struct {
    key    interface{}
    value  interface{}
    parent *node
    child  [2]*node
    b      int8
}

func (n *node) Key() interface{} {
    return n.key
}

func (n *node) Value() interface{} {
    return n.value
}

func (n *node) reset() {
    n.key = 0
    n.value = 0
    n.parent = nil
    for idx := range n.child {
        n.child[idx] = nil
    }
    n.b = 0
}

func (n *node) String() string {
    return fmt.Sprintf("%v", n.key)
}

func (n *node) Prev() *node {
    return n.walk1(0)
}

func (n *node) Next() *node {
    return n.walk1(1)
}

func (n *node) walk1(idx int) *node {
    if n == nil {
        return nil
    }

    tmp := n

    if tmp.child[idx] != nil {
        tmp := tmp.child[idx]
        /*
           case: 异或,相同为一,不同为零
        */
        for tmp.child[idx^1] != nil {
            tmp = tmp.child[idx^1]
        }
        return tmp
    }

    tmpParent := tmp.parent
    for tmpParent != nil && tmpParent.child[idx] == tmp {
        tmp = tmpParent
        tmpParent = tmpParent.parent
    }
    return tmpParent
}

func newNode() *node {
    return &node{}
}
