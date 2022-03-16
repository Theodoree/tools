package _000_1099



/*
1080. 根到叶路径上的不足节点
给定一棵二叉树的根 root，请你考虑它所有 从根到叶的路径：从根到任何叶的路径。（所谓一个叶子节点，就是一个没有子节点的节点）
假如通过节点 node 的每种可能的 “根-叶” 路径上值的总和全都小于给定的 limit，则该节点被称之为「不足节点」，需要被删除。
请你删除所有不足节点，并返回生成的二叉树的根。
*/
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	dummy := &TreeNode {
		Val:0,
		Left:root,
	}
	execute(root, dummy, 0, limit, 0)
	return dummy.Left
}

func execute(root *TreeNode, father *TreeNode, direction int, limit int, sum int) bool { //返回值是root点是否被删
	if nil == root {
		return false
	}
	sum += root.Val
	left := execute(root.Left, root, 0, limit, sum)
	right := execute(root.Right, root, 1, limit, sum)
	// 如果是叶子节点并且路径和小于limit 或 左子树被删并且没有右子树 或 右子树被删并且没有左子树 或左右子树都被删 则删掉root
	if (nil == root.Left && nil == root.Right && sum < limit) ||
		(left && nil == root.Right) || (right && nil == root.Left) || (left && right) {
		if 0 == direction {
			father.Left = nil
		} else {
			father.Right = nil
		}
		return true
	}
	return false
}

