package _00_299


/*
298. 二叉树最长连续序列
给你一棵指定的二叉树的根节点 root ，请你计算其中 最长连续序列路径 的长度。
最长连续序列路径 是依次递增 1 的路径。该路径，可以是从某个初始节点到树中任意节点，通过「父 - 子」关系连接而产生的任意路径。且必须从父节点到子节点，反过来是不可以的。
*/
func longestConsecutive(root *TreeNode) int {
	var target int
	dfs(root,0,0,&target)
	return target
}

func dfs(root *TreeNode,parent int,cur int,target *int){
	if root == nil {
		return
	}
	if root.Val-1 == parent {
		cur++
	}else{
		cur = 1
	}
	if cur > *target{
		*target = cur
	}
	dfs(root.Left,root.Val,cur,target)
	dfs(root.Right,root.Val,cur,target)
}
