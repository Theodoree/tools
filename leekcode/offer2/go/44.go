package _go

/*
剑指 Offer II 044. 二叉树每层的最大值
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
*/
func largestValues(root *TreeNode) []int {

	var result []int

	LargestValues(root, &result, 0)

	return result
}

func LargestValues(root *TreeNode, result *[]int, depth int) {
	if root == nil {
		return
	}

	if len(*result) <= depth {
		*result = append(*result, root.Val)
	} else {
		if t := (*result)[depth]; t < root.Val {

			(*result)[depth] = root.Val
		}
	}

	LargestValues(root.Left, result, depth+1)
	LargestValues(root.Right, result, depth+1)

}
