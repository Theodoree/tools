package _100_1199

type TreeNode struct {
	Val      int
	Children []*TreeNode
	Left     *TreeNode
	Right    *TreeNode
	Next     *TreeNode
}

/*
1120. 子树的最大平均值
给你一棵二叉树的根节点 root，找出这棵树的 每一棵 子树的 平均值 中的 最大 值。
子树是树中的任意节点和它的所有后代构成的集合。
树的平均值是树中节点值的总和除以节点数。
*/
func maximumAverageSubtree(root *TreeNode) float64 {
	_, _, avg := avgSubRoot(root)
	return avg
}

func avgSubRoot(root *TreeNode) (int, int, float64) {
	if root == nil {
		return 0, 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val, 1, float64(root.Val)
	}

	leftVal, leftNodeCnt, leftMaxAvgPrice := avgSubRoot(root.Left)
	rightVal, rightNodeCnt, rightMaxAvgPrice := avgSubRoot(root.Right)

	sumVal := leftVal + rightVal + root.Val
	sumNode := leftNodeCnt + rightNodeCnt + 1
	avgPrice := leftMaxAvgPrice
	if rightMaxAvgPrice > avgPrice {
		avgPrice = rightMaxAvgPrice
	}
	if float64(sumVal)/float64(sumNode) > avgPrice {
		avgPrice = float64(sumVal) / float64(sumNode)
	}

	return sumVal, sumNode, avgPrice
}
