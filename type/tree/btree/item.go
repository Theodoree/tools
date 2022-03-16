package btree

import "sort"

type items []Item

type Item interface {
    // Less tests whether the current item is less than the given argument.
    //
    // This must provide a strict weak ordering.
    // If !a.Less(b) && !b.Less(a), we treat this to mean a == b (i.e. we can only
    // hold one of either a or b in the tree).
    Less(than Item) bool
}

type ItemIterator func(i Item) bool

func (s *items) insertAt(index int, item Item) {
    *s = append(*s, nil)
    if index < len(*s) {
        copy((*s)[index+1:], (*s)[index:])
    }
    (*s)[index] = item
}

//删除指定下标item
func (s *items) removeAt(index int) Item {
    item := (*s)[index]
    copy((*s)[index:], (*s)[index+1:])
    (*s)[len(*s)-1] = nil // 因为不置为空,实际上data域还会有
    *s = (*s)[:len(*s)-1] // 设置length值
    return item
}

// pop
func (s *items) pop() (out Item) {
    index := len(*s) - 1
    out = (*s)[index]
    (*s)[index] = nil  // 因为不置为空,实际上data域还会有
    *s = (*s)[:index]
    return
}

func (s *items) truncate(index int) {
    var toClear items
    *s, toClear = (*s)[:index], (*s)[index:]
    for len(toClear) > 0 { // fixme:完全没看明白
        toClear = toClear[copy(toClear, nilItems):]
    }
}

func (s items) find(item Item) (index int, found bool) {
    i := sort.Search(len(s), func(i int) bool {
        return item.Less(s[i])
    })
    if i > 0 && !s[i-1].Less(item) {
        return i - 1, true
    }
    return i, false
}
