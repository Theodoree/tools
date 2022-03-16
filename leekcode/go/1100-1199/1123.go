package _100_1199

/*
1123. 最深叶节点的最近公共祖先
给你一个有根节点的二叉树，找到它最深的叶节点的最近公共祖先。
回想一下：
叶节点 是二叉树中没有子节点的节点
树的根节点的 深度 为 0，如果某一节点的深度为 d，那它的子节点的深度就是 d+1
如果我们假定 A 是一组节点 S 的 最近公共祖先，S 中的每个节点都在以 A 为根节点的子树中，且 A 的深度达到此条件下可能的最大值。
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
