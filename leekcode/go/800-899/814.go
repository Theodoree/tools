package _00_899

/*
814. 二叉树剪枝
给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。

返回移除了所有不包含 1 的子树的原二叉树。

节点 node 的子树为 node 本身加上所有 node 的后代。
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
