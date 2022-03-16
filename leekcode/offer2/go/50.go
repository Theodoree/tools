package _go

/*
剑指 Offer II 050. 向下的路径节点之和
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
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
