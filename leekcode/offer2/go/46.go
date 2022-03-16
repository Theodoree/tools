package _go

/*
剑指 Offer II 046. 二叉树的右侧视图
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/
func rightSideView(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var level []int
	LevelSort(root, &level, 0)

	return level

}

func LevelSort(tree *TreeNode, result *[]int, depth int) {

	if tree == nil {
		return
	}

	if len(*result) <= depth {
		*result = append(*result, 0)
	}
	(*result)[depth] = tree.Val
	LevelSort(tree.Left, result, depth+1)
	LevelSort(tree.Right, result, depth+1)
}
