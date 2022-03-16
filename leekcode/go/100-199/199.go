package _00_199

/*
199. 二叉树的右视图

给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
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
