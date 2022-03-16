package _100_1199



/*
1110. 删点成林
给出二叉树的根节点 root，树上每个节点都有一个不同的值。
如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
返回森林中的每棵树。你可以按任意顺序组织答案。
*/
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	del:=make([]bool,1001)
	var result []*TreeNode

	for _,v:=range to_delete{
		del[v] = true
	}
	if !del[root.Val]{
		result = append(result,root)
	}
	dfs(root,root,del,&result)
	return result
}

func dfs(root *TreeNode,parent *TreeNode,arr []bool,result *[]*TreeNode){
	if root == nil {
		return
	}

	dfs(root.Left,root,arr,result)
	dfs(root.Right,root,arr,result)

	if arr[root.Val] {
		if root == parent.Left{
			parent.Left = nil
		}
		if root == parent.Right{
			parent.Right = nil
		}
		if root.Left != nil {
			*result = append(*result,root.Left)
		}
		if root.Right != nil {
			*result = append(*result,root.Right)
		}

	}




}
