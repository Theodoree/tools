package _go

/*
剑指 Offer II 047. 二叉树剪枝
给定一个二叉树 根节点 root ，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点的值为 0 的子树。
节点 node 的子树为 node 本身，以及所有 node 的后代。
*/
func pruneTree(root *TreeNode) *TreeNode {

	PruneTree(root)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}

func PruneTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := PruneTree(root.Left)
	right := PruneTree(root.Right)

	if left == 0 {
		root.Left = nil
	}
	if right == 0 {
		root.Right = nil
	}

	return root.Val + left + right

}
