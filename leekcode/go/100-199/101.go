package _00_199

//对称二叉树
func isSymmetric(root *TreeNode) bool {
	return isSymmetri(root, root)

}

func isSymmetri(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSymmetri(q.Left, p.Right) && isSymmetri(p.Right, q.Left)
	}
	return false

}
