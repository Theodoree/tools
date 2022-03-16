package _go

/*
面试题 04.06. 后继者
设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
如果指定节点没有对应的“下一个”节点，则返回null。
*/
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	arr := [2]*TreeNode{}

	LDR(root, &arr, p)

	return arr[1]
}

func LDR(root *TreeNode, arr *[2]*TreeNode, p *TreeNode) {
	if root == nil {
		return
	}

	LDR(root.Left, arr, p)
	if arr[0] == p && arr[1] == nil {
		arr[1] = root
		return
	}
	arr[0] = root
	LDR(root.Right, arr, p)
}
