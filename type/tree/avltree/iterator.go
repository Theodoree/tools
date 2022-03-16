package avltree


// Iterator holding the iterator's state
type iterator struct {
    tree     *Tree
    node     *node
    position position
}

type position byte

const (
    begin, between, end position = 0, 1, 2
)

func (tree *Tree) Iterator() *iterator {
    return &iterator{tree: tree, node: nil, position: begin}
}

func (iterator *iterator) Next() bool {
    switch iterator.position {
    case begin:
        iterator.position = between
        iterator.node = iterator.tree.Left()
    case between:
        iterator.node = iterator.node.Next()
    }

    if iterator.node == nil {
        iterator.position = end
        return false
    }
    return true
}

func (iterator *iterator) Prev() bool {
    switch iterator.position {
    case end:
        iterator.position = between
        iterator.node = iterator.tree.Right()
    case between:
        iterator.node = iterator.node.Prev()
    }

    if iterator.node == nil {
        iterator.position = begin
        return false
    }
    return true
}

func (iterator *iterator) Value() interface{} {
    if iterator.node == nil {
        return nil
    }
    return iterator.node.value
}

func (iterator *iterator) Key() interface{} {
    if iterator.node == nil {
        return nil
    }
    return iterator.node.key
}

func (iterator *iterator) Begin() {
    iterator.node = nil
    iterator.position = begin
}

func (iterator *iterator) End() {
    iterator.node = nil
    iterator.position = end
}

func (iterator *iterator) First() bool {
    iterator.Begin()
    return iterator.Next()
}

func (iterator *iterator) Last() bool {
    iterator.End()
    return iterator.Prev()
}