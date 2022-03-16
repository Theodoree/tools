package _00_399


/*
308. 二维区域和检索 - 可变
给你一个二维矩阵 matrix ，你需要处理下面两种类型的若干次查询：
更新：更新 matrix 中某个单元的值
求和：计算矩阵 matrix 中某一矩形区域元素的 和 ，该区域由 左上角 (row1, col1) 和 右下角 (row2, col2) 界定。
*/
type NumMatrix struct {
	matrix [][]int
	preSum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	preSum:=make([][]int,len(matrix))
	for idx:=range preSum{
		preSum[idx] = make([]int,len(matrix[0]))
	}
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			if j == 0 {
				preSum[i][j] = matrix[i][j]
				continue
			}
			preSum[i][j] = preSum[i][j-1] + matrix[i][j]
		}
	}
	var n NumMatrix
	n.matrix = matrix
	n.preSum = preSum
	return n
}


func (this *NumMatrix) Update(row int, col int, val int)  {
	oldVal:=this.matrix[row][col]
	this.matrix[row][col] = val
	this.preSum[row][col] = this.preSum[row][col] -oldVal+val
	for j:=col+1;j<len(this.matrix[0]);j++{
		this.preSum[row][j] = this.preSum[row][j-1] + this.matrix[row][j]
	}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	for i:=row1;i<=row2;i++{
		if col1 == 0 {
			sum+=this.preSum[i][col2]
		}else{
			sum+=this.preSum[i][col2] - this.preSum[i][col1-1]
		}
	}
	return sum
}

