package _300_1399

/*
1315. 祖父节点值为偶数的节点和
给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：
该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）
如果不存在祖父节点值为偶数的节点，那么返回 0 。
*/
func init() {
	debug.SetGCPercent(0)
}
func sumEvenGrandparent(root *TreeNode) int {

	var sum int
	dfs(root.Left, root.Val, -1, &sum)
	dfs(root.Right, root.Val, -1, &sum)
	return sum
}

func dfs(root *TreeNode, parent, gradeFather int, result *int) {
	if root == nil {
		return
	}
	if gradeFather%2 == 0 {
		*result += root.Val
	}

	dfs(root.Left, root.Val, parent, result)
	dfs(root.Right, root.Val, parent, result)
}
