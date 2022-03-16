package _00_999




/*
931. 下降路径最小和
给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。
*/
func minFallingPathSum(matrix [][]int) int {

	min:= func(i,j int) int {
		if j< i {
			return j
		}
		return i
	}

	for i:=len(matrix)-2;i>=0;i--{

		for j:=0;j<len(matrix[i]);j++{
			if j == 0 {
				matrix[i][j] = matrix[i][j]+min(matrix[i+1][j],matrix[i+1][j+1])
			}else if j == len(matrix[i])-1{
				matrix[i][j] = matrix[i][j]+min(matrix[i+1][j],matrix[i+1][j-1])
			}else{
				matrix[i][j] = matrix[i][j]+min(matrix[i+1][j],min(matrix[i+1][j-1],matrix[i+1][j+1]))
			}
		}
	}

	result:=math.MaxInt64
	for i:=0;i<len(matrix);i++{
		if matrix[0][i] < result {
			result =  matrix[0][i]
		}
	}

	return result


}

