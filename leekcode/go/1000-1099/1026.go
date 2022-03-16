package _000_1099

/*
1026. 节点与其祖先之间的最大差值
给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。
（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
*/
func maxAncestorDiff(root *TreeNode) int {
	left := dfs(root.Left, root.Val, root.Val)
	right := dfs(root.Right, root.Val, root.Val)
	if left > right {
		return left
	}
	return right
}

func dfs(root *TreeNode, max, min int) int {
	if root == nil {
		return 0
	}

	if root.Val > max {
		max = root.Val
	}
	if root.Val < min {
		min = root.Val
	}
	if root.Left == nil && root.Right == nil {
		return max - min
	}

	left := dfs(root.Left, max, min)
	right := dfs(root.Right, max, min)

	if left > right {
		return left
	}
	return right
}
