package btree

import (
    "fmt"
    "io"
    "strings"
)

type node struct {
    items    items
    children children
    cow      *copyOnWriteContext
}


func (n *node) mutableFor(cow *copyOnWriteContext) *node {
    if n.cow == cow {
        return n
    }
    out := cow.newNode()
    if cap(out.items) >= len(n.items) {
        out.items = out.items[:len(n.items)]
    } else {
        out.items = make(items, len(n.items), cap(n.items))
    }
    copy(out.items, n.items)
    // Copy children
    if cap(out.children) >= len(n.children) {
        out.children = out.children[:len(n.children)]
    } else {
        out.children = make(children, len(n.children), cap(n.children))
    }
    copy(out.children, n.children)
    return out
}

func (n *node) mutableChild(i int) *node {
    c := n.children[i].mutableFor(n.cow)
    n.children[i] = c
    return c
}

func (n *node) split(i int) (Item, *node) {
    item := n.items[i]
    next := n.cow.newNode()
    next.items = append(next.items, n.items[i+1:]...)
    n.items.truncate(i)
    if len(n.children) > 0 {
        next.children = append(next.children, n.children[i+1:]...)
        n.children.truncate(i + 1)
    }
    return item, next
}

// maybeSplitChild checks if a child should be split, and if so splits it.
// Returns whether or not a split occurred.
func (n *node) maybeSplitChild(i, maxItems int) bool {
    if len(n.children[i].items) < maxItems {
        return false
    }
    first := n.mutableChild(i)
    item, second := first.split(maxItems / 2)
    n.items.insertAt(i, item)
    n.children.insertAt(i+1, second)
    return true
}


func (n *node) insert(item Item, maxItems int) Item {
    i, found := n.items.find(item)
    if found {
        out := n.items[i]
        n.items[i] = item
        return out
    }
    if len(n.children) == 0 {
        n.items.insertAt(i, item)
        return nil
    }
    if n.maybeSplitChild(i, maxItems) {
        inTree := n.items[i]
        switch {
        case item.Less(inTree):
            // no change, we want first split node
        case inTree.Less(item):
            i++ // we want second split node
        default:
            out := n.items[i]
            n.items[i] = item
            return out
        }
    }
    return n.mutableChild(i).insert(item, maxItems)
}

// get finds the given key in the subtree and returns it.
func (n *node) get(key Item) Item {
    i, found := n.items.find(key)
    if found {
        return n.items[i]
    } else if len(n.children) > 0 {
        return n.children[i].get(key)
    }
    return nil
}

func (n *node) remove(item Item, minItems int, typ toRemove) Item {
    var i int
    var found bool
    switch typ {
    case removeMax:
        if len(n.children) == 0 {
            return n.items.pop()
        }
        i = len(n.items)
    case removeMin:
        if len(n.children) == 0 {
            return n.items.removeAt(0)
        }
        i = 0
    case removeItem:
        i, found = n.items.find(item)
        if len(n.children) == 0 {
            if found {
                return n.items.removeAt(i)
            }
            return nil
        }
    default:
        panic("invalid type")
    }
    // If we get to here, we have children.
    if len(n.children[i].items) <= minItems {
        return n.growChildAndRemove(i, item, minItems, typ)
    }
    child := n.mutableChild(i)
    // Either we had enough items to begin with, or we've done some
    // merging/stealing, because we've got enough now and we're ready to return
    // stuff.
    if found {
        // The item exists at index 'i', and the child we've selected can give us a
        // predecessor, since if we've gotten here it's got > minItems items in it.
        out := n.items[i]
        // We use our special-case 'remove' call with typ=maxItem to pull the
        // predecessor of item i (the rightmost leaf of our immediate left child)
        // and set it into where we pulled the item from.
        n.items[i] = child.remove(nil, minItems, removeMax)
        return out
    }
    // Final recursive call.  Once we're here, we know that the item isn't in this
    // node and that the child is big enough to remove from.
    return child.remove(item, minItems, typ)
}

func (n *node) growChildAndRemove(i int, item Item, minItems int, typ toRemove) Item {
    if i > 0 && len(n.children[i-1].items) > minItems {
        // Steal from left child
        child := n.mutableChild(i)
        stealFrom := n.mutableChild(i - 1)
        stolenItem := stealFrom.items.pop()
        child.items.insertAt(0, n.items[i-1])
        n.items[i-1] = stolenItem
        if len(stealFrom.children) > 0 {
            child.children.insertAt(0, stealFrom.children.pop())
        }
    } else if i < len(n.items) && len(n.children[i+1].items) > minItems {
        // steal from right child
        child := n.mutableChild(i)
        stealFrom := n.mutableChild(i + 1)
        stolenItem := stealFrom.items.removeAt(0)
        child.items = append(child.items, n.items[i])
        n.items[i] = stolenItem
        if len(stealFrom.children) > 0 {
            child.children = append(child.children, stealFrom.children.removeAt(0))
        }
    } else {
        if i >= len(n.items) {
            i--
        }
        child := n.mutableChild(i)
        // merge with right child
        mergeItem := n.items.removeAt(i)
        mergeChild := n.children.removeAt(i + 1)
        child.items = append(child.items, mergeItem)
        child.items = append(child.items, mergeChild.items...)
        child.children = append(child.children, mergeChild.children...)
        n.cow.freeNode(mergeChild)
    }
    return n.remove(item, minItems, typ)
}

func (n *node) iterate(dir direction, start, stop Item, includeStart bool, hit bool, iter ItemIterator) (bool, bool) {
    var ok, found bool
    var index int
    switch dir {
    case ascend:
        if start != nil {
            index, _ = n.items.find(start)
        }
        for i := index; i < len(n.items); i++ {
            if len(n.children) > 0 {
                if hit, ok = n.children[i].iterate(dir, start, stop, includeStart, hit, iter); !ok {
                    return hit, false
                }
            }
            if !includeStart && !hit && start != nil && !start.Less(n.items[i]) {
                hit = true
                continue
            }
            hit = true
            if stop != nil && !n.items[i].Less(stop) {
                return hit, false
            }
            if !iter(n.items[i]) {
                return hit, false
            }
        }
        if len(n.children) > 0 {
            if hit, ok = n.children[len(n.children)-1].iterate(dir, start, stop, includeStart, hit, iter); !ok {
                return hit, false
            }
        }
    case descend:
        if start != nil {
            index, found = n.items.find(start)
            if !found {
                index = index - 1
            }
        } else {
            index = len(n.items) - 1
        }
        for i := index; i >= 0; i-- {
            if start != nil && !n.items[i].Less(start) {
                if !includeStart || hit || start.Less(n.items[i]) {
                    continue
                }
            }
            if len(n.children) > 0 {
                if hit, ok = n.children[i+1].iterate(dir, start, stop, includeStart, hit, iter); !ok {
                    return hit, false
                }
            }
            if stop != nil && !stop.Less(n.items[i]) {
                return hit, false //	continue
            }
            hit = true
            if !iter(n.items[i]) {
                return hit, false
            }
        }
        if len(n.children) > 0 {
            if hit, ok = n.children[0].iterate(dir, start, stop, includeStart, hit, iter); !ok {
                return hit, false
            }
        }
    }
    return hit, true
}

func (n *node) print(w io.Writer, level int) {
    fmt.Fprintf(w, "%sNODE:%v\n", strings.Repeat("  ", level), n.items)
    for _, c := range n.children {
        c.print(w, level+1)
    }
}

type children []*node


func (s *children) insertAt(index int, n *node) {
    *s = append(*s, nil)
    if index < len(*s) {
        copy((*s)[index+1:], (*s)[index:])
    }
    (*s)[index] = n
}

func (s *children) removeAt(index int) *node {
    n := (*s)[index]
    copy((*s)[index:], (*s)[index+1:])
    (*s)[len(*s)-1] = nil
    *s = (*s)[:len(*s)-1]
    return n
}

func (s *children) pop() (out *node) {
    index := len(*s) - 1
    out = (*s)[index]
    (*s)[index] = nil
    *s = (*s)[:index]
    return
}

func (s *children) truncate(index int) {
    var toClear children
    *s, toClear = (*s)[:index], (*s)[index:]
    for len(toClear) > 0 {
        toClear = toClear[copy(toClear, nilChildren):]
    }
}

