package _900_1999

type TreeNode struct {
	Val      int
	Children []*TreeNode
	Left     *TreeNode
	Right    *TreeNode
	Next     *TreeNode
}

/*
1973. 值等于子节点值之和的节点数量
给定一颗二叉树的根节点 root ，返回满足条件：节点的值等于该节点所有子节点的值之和 的节点的数量。
一个节点 x 的 子节点 是指从节点 x 出发，到所有叶子节点路径上的节点。没有子节点的节点的子节点和视为 0 。
*/
func equalToDescendants(root *TreeNode) int {
	var cnt int

	bfs(root, &cnt)

	return cnt
}

func bfs(root *TreeNode, cnt *int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			*cnt++
		}
		return root.Val
	}

	left := bfs(root.Left, cnt)
	right := bfs(root.Right, cnt)

	if left+right == root.Val {
		*cnt++
	}
	return left + right + root.Val
}
