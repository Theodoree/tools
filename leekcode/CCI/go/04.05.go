package _go

/*
面试题 04.05. 合法二叉搜索树
实现一个函数，检查一棵二叉树是否为二叉搜索树。
*/
func isValidBST(root *TreeNode) bool {
	return IsValidBST(root, math.MinInt64, math.MaxInt64)
}

func IsValidBST(root *TreeNode, min int, max int) bool {
	return root == nil || (root.Val > min && root.Val < max && IsValidBST(root.Left, min, root.Val) && IsValidBST(root.Right, root.Val, max))
}
