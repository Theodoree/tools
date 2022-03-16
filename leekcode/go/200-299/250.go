package _00_299

/*
250. 统计同值子树
给定一个二叉树，统计该二叉树数值相同的子树个数。
同值子树是指该子树的所有节点都拥有相同的数值。
*/
func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var c int
	CountUnivalSubtrees(root, &c)
	return c
}

func CountUnivalSubtrees(root *TreeNode, count *int) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		*count++
		return true
	}

	leftFlag := CountUnivalSubtrees(root.Left, count)
	rightFlag := CountUnivalSubtrees(root.Right, count)

	switch {
	case root.Left != nil && root.Right != nil:
		if leftFlag && rightFlag && root.Left.Val == root.Right.Val && root.Left.Val == root.Val {
			*count++
			return true
		}
	case root.Left != nil:
		if leftFlag && root.Left.Val == root.Val {
			*count++
			return true
		}
	case root.Right != nil:
		if rightFlag && root.Right.Val == root.Val {
			*count++
			return true
		}

	}

	return false

}
