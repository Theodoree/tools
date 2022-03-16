package _go

/*
面试题 04.12. 求和路径
给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。
示例:
给定如下二叉树，以及目标和 sum = 22，
*/

func pathSum(root *TreeNode, sum int) int {

	var cnt int
	PathSum(root, &cnt, sum)
	return cnt

}

func PathSum(root *TreeNode, cnt *int, sum int) {
	if root == nil {
		return
	}
	Sum(root, cnt, sum)
	PathSum(root.Left, cnt, sum)
	PathSum(root.Right, cnt, sum)
}

func Sum(root *TreeNode, cnt *int, sum int) {

	if root == nil {
		return
	}

	sum -= root.Val
	if sum == 0 {
		*cnt++
	}
	Sum(root.Left, cnt, sum)
	Sum(root.Right, cnt, sum)
}
