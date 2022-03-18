package _go


/*
剑指 Offer II 112. 最长递增路径
给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
对于每个单元格，你可以往上，下，左，右四个方向移动。 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
*/
func longestIncreasingPath(matrix [][]int) int {
	// 外层似乎是一个循环调用，这个maxVal每次都需要清理一下，不然上一个case的结果会影响当前case
	maxPath = 1
	// 初始化备忘录
	memo := make([][]int, len(matrix))
	for i := range memo{
		memo[i] = make([]int, len(matrix[0]))
	}
	// 遍历每一个结点进行递归
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			maxV := dfs(matrix, i, j, memo)
			// 记录每个结点的路径最大值
			maxPath = max(maxV, maxPath)
		}
	}
	return maxPath
}

var maxPath = 1

func dfs(matrix [][]int, i, j int, memo [][]int) int {
	if memo[i][j] != 0 { // 如果matrix[i][j]已经被设置了值，说明之前已经被求过最大路径了，直接返回即可
		return memo[i][j]
	}

	maxV := 1 // 最大值初始为1，当前自己算1个
	if i-1 >= 0 && matrix[i-1][j] > matrix[i][j] { // 校验下标和是否符合增长规律
		maxV = max(maxV, dfs(matrix, i-1, j, memo) + 1) // 是的话就递归下去，并将返回值+1（1代表自身）
	}

	if i+1 < len(matrix) && matrix[i+1][j] > matrix[i][j] {
		maxV = max(maxV, dfs(matrix, i+1, j, memo) + 1)
	}

	if j-1 >= 0 && matrix[i][j-1] > matrix[i][j] {
		maxV = max(maxV, dfs(matrix, i, j-1, memo) + 1)
	}

	if j+1 < len(matrix[0]) && matrix[i][j+1] > matrix[i][j] {
		maxV = max(maxV, dfs(matrix, i, j+1, memo) + 1)
	}
	// 上下左右都检查完，并且找出四个方向中路径最大值之后
	// 填到当前结点的备忘录中
	memo[i][j] = maxV
	// 返回当前结点的路径最大值，供上层使用
	return maxV
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
