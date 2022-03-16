package btree

type direction int

const (
	descend = direction(-1)
	ascend  = direction(+1)
)

var (
	nilItems    = make(items, 16)
	nilChildren = make(children, 16)
)

type toRemove int

const (
	removeItem toRemove = iota // removes the given item
	removeMin                  // removes smallest item in the subtree
	removeMax                  // removes largest item in the subtree
)

type freeType int

const (
	ftFreelistFull freeType = iota // node was freed (available for GC, not stored in freelist)
	ftStored                       // node was stored in the freelist for later use
	ftNotOwned                     // node was ignored by COW, since it's owned by another one
)
