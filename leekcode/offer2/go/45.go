package _go

/*
剑指 Offer II 045. 二叉树最底层最左边的值
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。
*/
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var result [2]int
	result[0] = 0
	result[1] = root.Val
	level(root, 0, &result)
	return result[1]
}

func level(root *TreeNode, depth int, result *[2]int) {
	if root == nil {
		return
	}

	if depth > result[0] {
		result[0] = depth
		result[1] = root.Val
	}
	level(root.Left, depth+1, result)
	level(root.Right, depth+1, result)

}
