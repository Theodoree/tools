package _00_599

/*
508. 出现次数最多的子树元素和
给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。
一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。
*/
func findFrequentTreeSum(root *TreeNode) []int {
	var m = map[int]int{}
	var max int
	dfs(root, m, &max)

	var arr []int
	for k, v := range m {
		if v == max {
			arr = append(arr, k)
		}
	}
	return arr
}
func dfs(root *TreeNode, m map[int]int, max *int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	left := dfs(root.Left, m, max)
	right := dfs(root.Right, m, max)
	val := m[root.Val+left+right]
	m[root.Val+left+right]++
	if val+1 > *max {
		*max = val + 1
	}

	return 0
}
