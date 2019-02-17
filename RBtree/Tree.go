package RBtree

type Tree struct {
	Root *Node
}

type RBtree interface {
	Insert(key int, value interface{})
	InsertNode(node *Node, key int)
	Find(key int) *Node
	Delete(key int)
	LeftRotate(node *Node)
	RightRotate(node *Node)
}

func (tree *Tree) Insert(key int, value interface{}) {

}

func (tree *Tree) InsertNode(node *Node, key int) {

}

func (tree *Tree) Find(key int) *Node {

}
func (tree *Tree) Delete(key int) {

}
func (tree *Tree) LeftRotate(node *Node) {

}
func (tree *Tree) RightRotate(node *Node) {

}
