package _go

/*
剑指 Offer 34. 二叉树中和为某一值的路径
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
*/
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	PathSum(root, targetSum, nil, &result)
	return result
}
func PathSum(root *TreeNode, targetSum int, cur []int, result *[][]int) {
	if root == nil {
		return
	}
	targetSum -= root.Val
	cur = append(cur, root.Val)

	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			*result = append(*result, cur)
		}
		return
	}

	PathSum(root.Left, targetSum, cur, result)
	buf := make([]int, len(cur))
	copy(buf, cur)
	PathSum(root.Right, targetSum, buf, result)
}

func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil { // 叶子节点
		if targetSum-root.Val == 0 { // 递归出口
			return [][]int{{root.Val}}
		} else {
			return nil
		}
	}
	var res [][]int
	if root.Left != nil {
		path := pathSum(root.Left, targetSum-root.Val)
		if path != nil {
			for _, v := range path {
				res = append(res, append([]int{root.Val}, v...))
			}
		}

	}

	if root != nil {
		path := pathSum(root.Right, targetSum-root.Val)
		if path != nil {
			for _, v := range path {
				res = append(res, append([]int{root.Val}, v...))
			}
		}
	}
	return res
}

func main() {
