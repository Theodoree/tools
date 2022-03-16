package _400_1499

/*
1469. 寻找所有的独生节点
二叉树中，如果一个节点是其父节点的唯一子节点，则称这样的节点为 “独生节点” 。二叉树的根节点不会是独生节点，因为它没有父节点。
给定一棵二叉树的根节点 root ，返回树中 所有的独生节点的值所构成的数组 。数组的顺序 不限 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLonelyNodes(root *TreeNode) []int {
	ans := []int{}
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil && r.Right == nil {
			ans = append(ans, r.Left.Val)
		} else if r.Left == nil && r.Right != nil {
			ans = append(ans, r.Right.Val)
		}
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return ans
}
