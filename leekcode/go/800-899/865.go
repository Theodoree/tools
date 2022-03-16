package _00_899

/*
865. 具有所有最深节点的最小子树
给定一个根为 root 的二叉树，每个节点的深度是 该节点到根的最短距离 。
如果一个节点在 整个树 的任意节点之间具有最大的深度，则该节点是 最深的 。
一个节点的 子树 是该节点加上它的所有后代的集合。
*/
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, node := dfs(root)
	return node
}

func dfs(r *TreeNode) (int, *TreeNode) {
	if r == nil {
		return 0, nil
	}

	left, ln := dfs(r.Left)
	right, rn := dfs(r.Right)
	if left > right {
		return left + 1, ln
	} else if left < right {
		return right + 1, rn
	}

	return left + 1, r
}
