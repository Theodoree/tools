package _go

import "runtime/debug"

/*
剑指 Offer II 053. 二叉搜索树中的中序后继
给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。
节点 p 的后继是值比 p.val 大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。
*/
func init() {
	debug.SetGCPercent(0)
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	arr := [2]*TreeNode{}

	LDR(root, &arr, p)

	return arr[1]
}

func LDR(root *TreeNode, arr *[2]*TreeNode, p *TreeNode) {
	if root == nil {
		return
	}

	LDR(root.Left, arr, p)
	if arr[0] == p && arr[1] == nil {
		arr[1] = root
		return
	}
	arr[0] = root
	LDR(root.Right, arr, p)
}
