package _000_1099

/*
1038. 把二叉搜索树转换为累加树
给定一个二叉搜索树，请将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。
*/

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	f(root, sum)
	return root
}

func f(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	sum = f(root.Right, sum)
	root.Val += sum
	sum = root.Val
	sum = f(root.Left, sum)
	return sum
}
