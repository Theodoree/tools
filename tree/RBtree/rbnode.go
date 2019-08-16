package RBtree

type RBNode struct {
	Parent      *RBNode
	Key         int64
	Color       bool
	Value       int64
	Left, Right *RBNode
}

func (n *RBNode) GetParent() *RBNode {
	return n.Parent
}

func (n *RBNode) GetGrandParent() *RBNode {
	return n.Parent.Parent
}

func (n *RBNode) GetUncle() (*RBNode, string) {
	parent := n.GetParent()
	grandParent := n.GetGrandParent()

	if grandParent.Left == parent {
		return grandParent.Right, "right"
	}
	return grandParent.Left, "left"
}

func (n *RBNode) IsRed() bool {
	return n.Color == RED
}

func (n *RBNode) IsBlack() bool {
	return n.Color == BLACK

}

func (n *RBNode) SetRed() {
	n.Color = RED
}

func (n *RBNode) SetBlack() {
	n.Color = BLACK

}
