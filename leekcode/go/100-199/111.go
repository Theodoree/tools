package _00_199

/*
111. 二叉树的最小深度

给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	t := -1
	if root.Left != nil {
		t = minDepth(root.Left)
	}

	if root.Right != nil {
		if t == -1 {
			t = minDepth(root.Right)
		} else {
			t = min(t, minDepth(root.Right))
		}
	}

	if t == -1 {
		return 1
	}

	return 1 + t
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
