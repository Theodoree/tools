package _00_399

/*
304. 二维区域和检索 - 矩阵不可变
给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。
*/

type NumMatrix struct {
	preSum []int
	l      int
}

func Constructor(matrix [][]int) NumMatrix {
	var c NumMatrix
	c.preSum = make([]int, len(matrix)*len(matrix[0])+1)
	var count int
	count = 1
	for _, v := range matrix {
		for _, j := range v {
			if count > 0 {
				c.preSum[count] += c.preSum[count-1] + j
			} else {
				c.preSum[count] = j
			}
			count++
		}
	}
	c.l = len(matrix[0])
	return c
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	for i := row1; i <= row2; i++ {
		if i*this.l+col1-1 == -1 {
			sum += this.preSum[i*this.l+col2+1]
		} else {
			sum += this.preSum[i*this.l+col2+1] - this.preSum[i*this.l+col1]
		}
	}
	return sum
}
