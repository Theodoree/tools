package _go

/*
剑指 Offer 32 - III. 从上到下打印二叉树 III
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var arr [2][]*TreeNode
	var flag bool
	var depth int
	arr[0] = append(arr[0], root)

	for len(arr[0]) > 0 {
		result = append(result, []int{})
		for idx, v := range arr[0] {
			if v == nil {
				continue
			}
			if flag {
				result[depth] = append(result[depth], arr[0][len(arr[0])-idx-1].Val)
			} else {
				result[depth] = append(result[depth], v.Val)
			}
			if v.Left != nil {
				arr[1] = append(arr[1], v.Left)
			}
			if v.Right != nil {
				arr[1] = append(arr[1], v.Right)
			}
		}
		arr[0], arr[1] = arr[1], arr[0]
		arr[1] = arr[1][:0]
		flag = !flag
		depth++
	}

	return result

}
