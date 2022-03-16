package _00_399

/*
333. 最大 BST 子树
给定一个二叉树，找到其中最大的二叉搜索树（BST）子树，并返回该子树的大小。其中，最大指的是子树节点数最多的。

二叉搜索树（BST）中的所有节点都具备以下属性：

左子树的值小于其父（根）节点的值。

右子树的值大于其父（根）节点的值。

注意：子树必须包含其所有后代。
*/

func largestBSTSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	target := 1
	dfs(root, -1e9, 1e9, &target)

	return target
}

func dfs(root *TreeNode, max int, min int, target *int) (bool, int, int, int) {
	if root == nil {
		return true, 0, 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return true, 1, root.Val, root.Val
	}

	ok1, left, leftMax, leftMin := dfs(root.Left, max, min, target)
	ok2, right, rightMax, rightMin := dfs(root.Right, max, min, target)
	if !(ok1 && ok2) || (left > 0 && root.Val <= leftMax) || (right > 0 && root.Val >= rightMin) {
		return false, 0, 0, 0
	}

	if root.Val > max {
		max = root.Val
	}
	if min > root.Val {
		min = root.Val
	}

	if left > 0 {
		if leftMax > max {
			max = leftMax
		}
		if leftMin < min {
			min = leftMin
		}
	}
	if right > 0 {
		if rightMax > max {
			max = rightMax
		}
		if rightMin < min {
			min = rightMin
		}
	}
	sum := left + right + 1
	if *target < sum {
		*target = sum
	}
	return true, sum, max, min
}
