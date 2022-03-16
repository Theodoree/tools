package _00_199

/*
108. 将有序数组转换为二叉搜索树

将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-5,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	n := SortedArrayToBST(nums, 0, len(nums)-1)
	return n
}

func SortedArrayToBST(nums []int, l, r int) *TreeNode {

	if l > r {
		return nil
	}

	m := l + (r-l)/2
	node := &TreeNode{Val: nums[m]}
	node.Left = SortedArrayToBST(nums, l, m-1)
	node.Right = SortedArrayToBST(nums, m+1, r)
	return node
}
