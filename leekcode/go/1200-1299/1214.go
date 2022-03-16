package _200_1299


type TreeNode struct {
	Val      int
	Left     *TreeNode
	Right    *TreeNode
}


/*
1214. 查找两棵二叉搜索树之和
给出两棵二叉搜索树的根节点 root1 和 root2 ，请你从两棵树中各找出一个节点，使得这两个节点的值之和等于目标值 Target。
如果可以找到返回 True，否则返回 False。
*/
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	var arr []int
	DLR(root1,&arr)
	var m =make(map[int]struct{})
	for _,v:=range arr{
		m[v] = struct{}{}
	}

	arr = arr[:0]
	DLR(root2,&arr)
	for _,v:=range arr {
		if _,ok:=m[target-v];ok{
			return true
		}
	}

	return false
}

func DLR(root *TreeNode,result *[]int){
	if root == nil {
		return
	}

	*result = append(*result,root.Val)
	DLR(root.Left,result)
	DLR(root.Right,result)
}
