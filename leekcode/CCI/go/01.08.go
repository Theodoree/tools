package _go

import "math"

/*
面试题 01.08. 零矩阵
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
*/
func setZeroes(matrix [][]int) {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				// 上
				curi := i - 1
				for curi >= 0 && matrix[curi][j] != 0 {
					matrix[curi][j] = math.MinInt64
					curi--
				}

				// 下
				curi = i + 1
				for curi < len(matrix) && matrix[curi][j] != 0 {
					matrix[curi][j] = math.MinInt64
					curi++
				}

				// 左
				curJ := j - 1
				for curJ >= 0 && matrix[i][curJ] != 0 {
					matrix[i][curJ] = math.MinInt64
					curJ--
				}

				// 右
				curJ = j + 1
				for curJ < len(matrix[i]) && matrix[i][curJ] != 0 {
					matrix[i][curJ] = math.MinInt64
					curJ++
				}

			}

		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == math.MinInt64 {
				matrix[i][j] = 0
			}

		}
	}

}
