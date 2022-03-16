package _100_1199


/*
1161. 最大层内元素和
给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
请返回层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。
*/
func maxLevelSum(root *TreeNode) int {

	var result []int
	dfs(root,0,&result)

	var max int
	max = math.MinInt64
	var idx int
	for i:=0;i<len(result);i++{
		if result[i] > max {
			max = result[i]
			idx =i
		}
	}

	return idx+1
}

func dfs(root *TreeNode,depth int,result *[]int){
	if root == nil {
		return
	}
	if len(*result) <= depth{
		*result = append(*result,0)
	}
	(*result)[depth]+=root.Val
	dfs(root.Left,depth+1,result)
	dfs(root.Right,depth+1,result)
}

