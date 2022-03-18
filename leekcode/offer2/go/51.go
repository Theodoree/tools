package _go


func maxPathSum(root *TreeNode) int {
	var max int
	max = math.MinInt32
	DFS(root, &max)
	return max
}

func DFS(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := DFS(root.Left, max)
	right := DFS(root.Right, max)
	lmr := root.Val + GetMax(0, left) + GetMax(0, right)
	ret := root.Val + GetMax(0, GetMax(left, right))
	*max = GetMax(*max, GetMax(lmr, ret))

	return ret
}

func GetMax(nums ...int) int {
	var Max int
	Max = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > Max {
			Max = nums[i]
		}

	}
	return Max
}

