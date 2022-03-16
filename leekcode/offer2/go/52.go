package _go

/*
剑指 Offer II 052. 展平二叉搜索树
给你一棵二叉搜索树，请 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
*/
func increasingBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	head := new(TreeNode)
	result := head
	IncreasingBST(root, result)

	return head.Right
}

func IncreasingBST(root *TreeNode, result *TreeNode) *TreeNode {
	if root == nil {
		return result
	}
	result = IncreasingBST(root.Left, result)
	result.Right = &TreeNode{Val: root.Val}
	result = result.Right
	result = IncreasingBST(root.Right, result)
	return result

}
