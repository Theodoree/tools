package __99



func inorderTraversal(root *TreeNode) []int {

	var result []int
	ad(root, &result)

	return result
}

func ad(root *TreeNode, result *[]int) {
	if root != nil {
		ad(root.Left, result) // L
		*result = append(*result,root.Val)
		ad(root.Right, result) // R
	}
}