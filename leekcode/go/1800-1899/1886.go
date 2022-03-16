package _800_1899

/*
1886. 判断矩阵经轮转后是否一致
给你两个大小为 n x n 的二进制矩阵 mat 和 target 。现 以 90 度顺时针轮转 矩阵 mat 中的元素 若干次 ，如果能够使 mat 与 target 一致，返回 true ；否则，返回 false 。
*/
func findRotation(mat [][]int, target [][]int) bool {
	var flag = [4]bool{true, true, true, true}
	n := len(mat[0]) - 1

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			// 1 : a[i][j] -> a[j][col-i]
			// 2 : a[i][j] -> a[row-i][col-j]
			// 3 : a[i][j] -> a[col-j][i]
			// 4 : a[i][j] -> a[i][j]

			// 90
			if mat[i][j] != target[j][n-i] {
				flag[0] = false
			}

			// 180
			if mat[i][j] != target[n-i][n-j] {
				flag[1] = false
			}

			// 240
			if mat[i][j] != target[n-j][i] {
				flag[2] = false
			}

			// 360
			if mat[i][j] != target[i][j] {
				flag[3] = false
			}
		}
	}

	if flag[0] || flag[1] || flag[2] || flag[3] {
		return true
	}
	return false
}
