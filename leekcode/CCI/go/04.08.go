package _go

/*
面试题 04.08. 首个共同祖先
设计并实现一个算法，找出二叉树中某两个节点的第一个共同祖先。不得将其他的节点存储在另外的数据结构中。注意：这不一定是二叉搜索树。
例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return nil
		// 如果left和right都不为空,则说明是最近的公共节点
	} else if left != nil && right != nil {
		return root
	} else {
		// 如果有一个为空,则继续往上算
		if left == nil {
			return right
		}

		return left
	}
}
