package _go

import "runtime/debug"

/*
剑指 Offer 26. 树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。
*/
func init() {
	debug.SetGCPercent(0)
}
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return IsSubStructure(A, B, false)
}
func IsSubStructure(A *TreeNode, B *TreeNode, cmp bool) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val == B.Val {
		if IsSubStructure(A.Left, B.Left, true) && IsSubStructure(A.Right, B.Right, true) {
			return true
		}
	}
	if cmp {
		return false
	}

	return IsSubStructure(A.Left, B, false) || IsSubStructure(A.Right, B, false)
}
