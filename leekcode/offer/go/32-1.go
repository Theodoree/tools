package _go

/*
剑指 Offer 32 - I. 从上到下打印二叉树
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
*/

func levelOrder(root *TreeNode) []int {
	var result []int
	var arr [2][]*TreeNode
	arr[0] = append(arr[0], root)

	for len(arr[0]) > 0 {
		for _, v := range arr[0] {
			if v == nil {
				continue
			}
			result = append(result, v.Val)
			arr[1] = append(arr[1], v.Left)
			arr[1] = append(arr[1], v.Right)
		}
		arr[0], arr[1] = arr[1], arr[0]
		arr[1] = arr[1][:0]
	}

	return result

}
