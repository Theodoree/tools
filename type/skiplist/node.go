package skiplist

type node struct {
    forward    []*node // 前驱
    backward   *node   // 后继
    key, value interface{}
}

// 后继
func (n *node) next() *node {
    if len(n.forward) == 0 {
        return nil
    }
    return n.forward[0]
}

// 前驱节点
func (n *node) previous() *node {
    return n.backward
}

func (n *node) hasNext() bool {
    return n.next() != nil
}

func (n *node) hasPrevious() bool {
    return n.previous() != nil
}
