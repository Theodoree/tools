package _400_1499

/*
1448. 统计二叉树中好节点的数目
给你一棵根为 root 的二叉树，请你返回二叉树中好节点的数目。
「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。
*/
func goodNodes(root *TreeNode) int {
	count:=0
	goodNodes1(root,math.MinInt64,&count)
	return count
}

func goodNodes1(root *TreeNode,max int,count *int) {
	if root == nil {
		return
	}

	if root.Val >= max {
		*count++
		max = root.Val
	}

	goodNodes1(root.Left,max,count)
	goodNodes1(root.Right,max,count)




}
