package _go

/*
面试题 04.04. 检查平衡性
实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。
*/
func isBalanced(root *TreeNode) bool {

	return root == nil || math.Abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func max(i, j float64) float64 {
	if i > j {
		return i
	}
	return j

}
func depth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}
