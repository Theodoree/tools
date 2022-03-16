package bts

type BTS struct {
	root *Node
}
type Node struct {
	Left  *Node
	Right *Node
	Key   int
	Val   int
}

/*
	性质1: 左孩子必定大于父节点
	性质2: 右孩子必定小于父节点
*/

func NewBTS() *BTS {
	return &BTS{}
}

func (t *BTS) Delete(key int) (bool, int) {
	if t.root == nil {
		return false, 0
	}
	if t.root.key == key { // delete root

	}

}

func (t *BTS) Find(key int) (*Node, *Node) {
	if t.root == nil {
		return nil, nil
	}

}

func (t *BTS) Insert(key int, val int) {
	if t.root == nil {
		t.root = newNode(key, val)
		return
	}

	cur := t.root

}
func newNode(key, val int) *Node {
	return &Node{key: key, val: val}
}
