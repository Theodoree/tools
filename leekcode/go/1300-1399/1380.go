package _300_1399

/*
1380. 矩阵中的幸运数
给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。

幸运数是指矩阵中满足同时下列两个条件的元素：

在同一行的所有元素中最小
在同一列的所有元素中最大


示例 1：

输入：matrix = [[3,7,8],[9,11,13],[15,16,17]]
输出：[15]
解释：15 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
示例 2：

输入：matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
输出：[12]
解释：12 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
示例 3：

输入：matrix = [[7,8],[1,2]]
输出：[7]
*/
func luckyNumbers(matrix [][]int) []int {

	var result []int

	maxByCol := make([]int, len(matrix[0]))
	minByRow := make([]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		minByRow[i] = math.MaxInt64
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < minByRow[i] {
				minByRow[i] = matrix[i][j]
			}
			if matrix[i][j] > maxByCol[j] {
				maxByCol[j] = matrix[i][j]
			}
		}
	}

	for _, v := range maxByCol {
		for i := 0; i < len(minByRow); i++ {
			if v == minByRow[i] {
				result = append(result, v)
			}
		}
	}

	return result
}
