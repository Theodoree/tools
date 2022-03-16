package _400_1499



/*
1457. 二叉树中的伪回文路径
给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。
请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。
*/
func pseudoPalindromicPaths (root *TreeNode) int {
	return dfs(root,0,)
}



func dfs(root *TreeNode,result int) int {
	if root == nil {
		return 0
	}
	result^= 1<< root.Val
	if root.Left == nil && root.Right == nil {
		var flag bool
		for i:=0;i<=9 ;i++{
			if (1 << i)&result == 0 {
				continue
			}
			if flag{
				return 0
			}
			flag = true
		}
		return 1
	}
	return dfs(root.Left,result) + dfs(root.Right,result)
}
