package _600_1699



/*
1650. 二叉树的最近公共祖先 III
给定一棵二叉树中的两个节点 p 和 q，返回它们的最近公共祖先节点（LCA）。
每个节点都包含其父节点的引用（指针）。Node 的定义如下：
*/

type Node struct {
	Val int
	Left *Node
	Right *Node
	Parent *Node
}
func lowestCommonAncestor(p *Node, q *Node) *Node {
	root:=p
	for root.Parent != nil {
		root = root.Parent
	}

	return lowestCommonAncestor1(root,p,q)

}

func lowestCommonAncestor1(root, p, q *Node) *Node {

	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	if left == nil && right == nil {
		return nil
		// 如果left和right都不为空,则说明是最近的公共节点
	} else if left != nil && right != nil {
		return root
	} else {
		// 如果有一个为空,则继续往上算
		if left == nil {
			return right
		}

		return left
	}
}
