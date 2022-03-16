package _00_699

/*
663. 均匀树划分
给定一棵有 n 个结点的二叉树，你的任务是检查是否可以通过去掉树上的一条边将树分成两棵，且这两棵树结点之和相等。
*/

func checkEqualTree(root *TreeNode) bool {

	var ans bool
	dfs(root, sum(root), &ans)
	return ans
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sum(root.Left) + sum(root.Right) + root.Val
}

func dfs(root *TreeNode, sum int, ans *bool) int {
	if root == nil || *ans {
		return 0
	}

	ls := dfs(root.Left, sum, ans)
	rs := dfs(root.Right, sum, ans)
	if (root.Left != nil && ls == sum-ls) || (root.Right != nil && rs == sum-rs) {
		*ans = true
	}

	return ls + rs + root.Val
}
