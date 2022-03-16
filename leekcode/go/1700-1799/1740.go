package _700_1799

/*
1740. 找到二叉树中的距离
给定一棵二叉树的根节点 root 以及两个整数 p 和 q ，返回该二叉树中值为 p 的结点与值为 q 的结点间的 距离 。
两个结点间的 距离 就是从一个结点到另一个结点的路径上边的数目。
*/

func findDistance(root *TreeNode, p int, q int) int {
	if p == q {
		return 0
	}

	ptr := findParent(root, p, q)

	return getDistance(ptr, p) + getDistance(ptr, q)
}

func findParent(root *TreeNode, p int, q int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p || root.Val == q {
		return root
	}
	left := findParent(root.Left, p, q)
	right := findParent(root.Right, p, q)
	if left != nil && right != nil { // 公共子节点
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func getDistance(root *TreeNode, x int) int {
	if root == nil {
		return -1
	}
	if root.Val == x {
		return 0
	}
	l := getDistance(root.Left, x)
	if l != -1 {
		return l + 1
	}
	r := getDistance(root.Right, x)
	if r != -1 {
		return r + 1
	}
	return -1
}
