package _go



/*
剑指 Offer II 013. 二维子矩阵的和
给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回左上角 (row1, col1) 、右下角 (row2, col2) 的子矩阵的元素总和。
*/

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix))
	// tmp := make([]int, len(matrix[0]))
	for i := range sum {
		sum[i] = make([]int, len(matrix[0]))

		for j := range sum[i] {
			sum[i][j] = matrix[i][j]
			if j - 1 >= 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i - 1 >= 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j - 1 >= 0 && i - 1 >= 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}
	return NumMatrix{sum}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := this.sum[row2][col2]
	row1--
	col1--
	if row1 >= 0 {
		res -= this.sum[row1][col2]
	}
	if col1 >= 0 {
		res -= this.sum[row2][col1]
	}
	if row1 >= 0 && col1 >= 0 {
		res += this.sum[row1][col1]
	}
	return res
}
