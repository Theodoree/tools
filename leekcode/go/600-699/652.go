package _00_699

import "strconv"

/*
652. 寻找重复的子树
给定一棵二叉树 root，返回所有重复的子树。

对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

如果两棵树具有相同的结构和相同的结点值，则它们是重复的。
*/
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var m = map[string]int{}
	var result []*TreeNode
	dfs(root,&result,m)
	return result
}
func dfs(root *TreeNode,result *[]*TreeNode,m map[string]int) string{
	if root == nil {
		return "#"
	}
	str:=strconv.Itoa(root.Val) + " " + dfs(root.Left,result,m) + " " + dfs(root.Right,result,m)
	if m[str] == 1 {
		*result = append(*result,root)
	}
	m[str]++
	return str
}

