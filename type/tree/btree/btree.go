package btree

type BTree struct {
    degree int
    length int
    root   *node
    cow    *copyOnWriteContext
}

func New(degree int) *BTree {
    return NewWithFreeList(degree, NewFreeList())
}


func (t *BTree) Clone() (t2 *BTree) {
    // Create two entirely new copy-on-write contexts.
    // This operation effectively creates three trees:
    //   the original, shared nodes (old b.cow)
    //   the new b.cow nodes
    //   the new out.cow nodes
    cow1, cow2 := *t.cow, *t.cow
    out := *t
    t.cow = &cow1
    out.cow = &cow2
    return &out
}

func (t *BTree) maxItems() int {
    return t.degree*2 - 1
}

func (t *BTree) minItems() int {
    return t.degree - 1
}

type copyOnWriteContext struct {
    freelist *FreeList
}

func (c *copyOnWriteContext) newNode() (n *node) {
    n = c.freelist.newNode()
    n.cow = c
    return
}

func (c *copyOnWriteContext) freeNode(n *node) freeType {
    if n.cow == c {
        // clear to allow GC
        n.items.truncate(0)
        n.children.truncate(0)
        n.cow = nil
        c.freelist.freeNode(n)
        return ftStored
    } else {
        return ftNotOwned
    }
}

func NewWithFreeList(degree int, f *FreeList) *BTree {
    if degree <= 1 {
        panic("bad degree")
    }
    return &BTree{
        degree: degree,
        cow:    &copyOnWriteContext{freelist: f},
    }
}

func (t *BTree) ReplaceOrInsert(item Item) Item {
    if item == nil {
        panic("nil item being added to BTree")
    }
    if t.root == nil {
        t.root = t.cow.newNode()
        t.root.items = append(t.root.items, item)
        t.length++
        return nil
    } else {
        t.root = t.root.mutableFor(t.cow)
        if len(t.root.items) >= t.maxItems() {
            item2, second := t.root.split(t.maxItems() / 2)
            oldroot := t.root
            t.root = t.cow.newNode()
            t.root.items = append(t.root.items, item2)
            t.root.children = append(t.root.children, oldroot, second)
        }
    }
    out := t.root.insert(item, t.maxItems())
    if out == nil {
        t.length++
    }
    return out
}

func (t *BTree) Delete(item Item) Item {
    return t.deleteItem(item, removeItem)
}

func (t *BTree) DeleteMin() Item {
    return t.deleteItem(nil, removeMin)
}

func (t *BTree) DeleteMax() Item {
    return t.deleteItem(nil, removeMax)
}

func (t *BTree) deleteItem(item Item, typ toRemove) Item {
    if t.root == nil || len(t.root.items) == 0 {
        return nil
    }
    t.root = t.root.mutableFor(t.cow)
    out := t.root.remove(item, t.minItems(), typ)
    if len(t.root.items) == 0 && len(t.root.children) > 0 {
        oldroot := t.root
        t.root = t.root.children[0]
        t.cow.freeNode(oldroot)
    }
    if out != nil {
        t.length--
    }
    return out
}

func (t *BTree) AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator) {
    if t.root == nil {
        return
    }
    t.root.iterate(ascend, greaterOrEqual, lessThan, true, false, iterator)
}

func (t *BTree) AscendLessThan(pivot Item, iterator ItemIterator) {
    if t.root == nil {
        return
    }
    t.root.iterate(ascend, nil, pivot, false, false, iterator)
}

func (t *BTree) AscendGreaterOrEqual(pivot Item, iterator ItemIterator) {
    if t.root == nil {
        return
    }
    t.root.iterate(ascend, pivot, nil, true, false, iterator)
}

func (t *BTree) Ascend(iterator ItemIterator) {
    if t.root == nil {
        return
    }
    t.root.iterate(ascend, nil, nil, false, false, iterator)
}

func (t *BTree) DescendRange(lessOrEqual, greaterThan Item, iterator ItemIterator) {
    if t.root == nil {
        return
    }
    t.root.iterate(descend, lessOrEqual, greaterThan, true, false, iterator)
}

func (t *BTree) DescendLessOrEqual(pivot Item, iterator ItemIterator) {
    if t.root == nil {
        return
    }
    t.root.iterate(descend, pivot, nil, true, false, iterator)
}

func (t *BTree) DescendGreaterThan(pivot Item, iterator ItemIterator) {
    if t.root == nil {
        return
    }
    t.root.iterate(descend, nil, pivot, false, false, iterator)
}

func (t *BTree) Descend(iterator ItemIterator) {
    if t.root == nil {
        return
    }
    t.root.iterate(descend, nil, nil, false, false, iterator)
}

func (t *BTree) Get(key Item) Item {
    if t.root == nil {
        return nil
    }
    return t.root.get(key)
}

func (t *BTree) Min() Item {
    return min(t.root)
}

func (t *BTree) Max() Item {
    return max(t.root)
}

func (t *BTree) Has(key Item) bool {
    return t.Get(key) != nil
}

func (t *BTree) Len() int {
    return t.length
}

func (t *BTree) Clear(addNodesToFreelist bool) {
    if t.root != nil && addNodesToFreelist {
        t.root.reset(t.cow)
    }
    t.root, t.length = nil, 0
}

func (n *node) reset(c *copyOnWriteContext) bool {
    for _, child := range n.children {
        if !child.reset(c) {
            return false
        }
    }
    return c.freeNode(n) != ftFreelistFull
}
