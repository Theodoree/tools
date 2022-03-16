package _500_1599

/*
1582. 二进制矩阵中的特殊位置
给你一个大小为 rows x cols 的矩阵 mat，其中 mat[i][j] 是 0 或 1，请返回 矩阵 mat 中特殊位置的数目 。
特殊位置 定义：如果 mat[i][j] == 1 并且第 i 行和第 j 列中的所有其他元素均为 0（行和列的下标均 从 0 开始 ），则位置 (i, j) 被称为特殊位置。
*/
func numSpecial(mat [][]int) int {
	var count int
	for i := 0; i < len(mat); i++ {
		numIdx := -1
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				if numIdx != -1 {
					numIdx = -1
					break
				}
				numIdx = j
			}
		}

		if numIdx != -1 { // 检查列
			up := i - 1
			down := i + 1
			flag := true
			for (up >= 0 || down < len(mat)) && flag {
				if up >= 0 && mat[up][numIdx] != 0 {
					flag = false
				}

				if down < len(mat) && mat[down][numIdx] != 0 {
					flag = false
				}
				up--
				down++
			}
			if flag {
				count++
			}
		}
	}

	return count
}
