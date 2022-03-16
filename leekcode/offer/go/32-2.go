package _go

/*
剑指 Offer 32 - II. 从上到下打印二叉树 II
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
*/
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	dfs(root, 0, &result)
	return result
}

func dfs(root *TreeNode, depth int, result *[][]int) {
	if root == nil {
		return
	}

	if len(*result) == depth {
		*result = append(*result, []int{root.Val})
	} else {
		v := *result
		v[depth] = append(v[depth], root.Val)
	}

	dfs(root.Left, depth+1, result)
	dfs(root.Right, depth+1, result)

}
