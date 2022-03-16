package _00_499

//450. 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		minNode, minNodeParent := min(root.Right)
		minNodeParent.Left = nil
		minNode.Left = root.Left
		minNode.Right = root.Right

		return minNode
	}

	if root.Val < key {
		root.Right = deleteNode(root.Left, key)
		return root
	}

	if root.Val > key {
		root.Left = deleteNode(root.Right, key)

		return root
	}

	return nil
}

func min(n *TreeNode) (*TreeNode, *TreeNode) {
	current := n
	var parten *TreeNode
	for {
		if current.Left != nil {
			parten = current
			current = current.Left
		} else {
			return current, parten
		}
	}
}
