package _400_1499


/*
1485. 克隆含随机指针的二叉树
给你一个二叉树，树中每个节点都含有一个附加的随机指针，该指针可以指向树中的任何节点或者指向空（null）。

请返回该树的 深拷贝 。

该树的输入/输出形式与普通二叉树相同，每个节点都用 [val, random_index] 表示：

val：表示 Node.val 的整数
random_index：随机指针指向的节点（在输入的树数组中）的下标；如果未指向任何节点，则为 null 。
该树以 Node 类的形式给出，而你需要以 NodeCopy 类的形式返回克隆得到的树。NodeCopy 类和Node 类定义一致。
*/

func copyRandomBinaryTree(root *Node) *NodeCopy {
	if root == nil {
		return nil
	}

	table := make(map[uintptr]*NodeCopy)
	c:=root.Clone(table)
	root.CloneRandom(c,table)
	return c
}

func (n *Node) Clone( table map[uintptr]*NodeCopy) *NodeCopy {
	var cur = &NodeCopy{
		Val: n.Val,
	}
	table[uintptr(unsafe.Pointer(n))] = cur
	if n.Left != nil {
		cur.Left = n.Left.Clone(table)
	}
	if n.Right != nil {
		cur.Right = n.Right.Clone(table)
	}

	return cur
}

func (n *Node) CloneRandom( head *NodeCopy,table map[uintptr]*NodeCopy){
	if head.Left != nil {
		n.Left.CloneRandom(head.Left,table)
	}
	if head.Right != nil {
		n.Right.CloneRandom(head.Right,table)
	}
	if n.Random != nil {
		head.Random = table[uintptr(unsafe.Pointer(n.Random))]
	}
}

