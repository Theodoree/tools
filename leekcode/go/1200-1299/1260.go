package _200_1299

/*
1260. 二维网格迁移
给你一个 m 行 n 列的二维网格 grid 和一个整数 k。你需要将 grid 迁移 k 次。
每次「迁移」操作将会引发下述活动：

位于 grid[i][j] 的元素将会移动到 grid[i][j + 1]。
位于 grid[i][n - 1] 的元素将会移动到 grid[i + 1][0]。
位于 grid[m - 1][n - 1] 的元素将会移动到 grid[0][0]。
请你返回 k 次迁移操作后最终得到的 二维网格。
*/
func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid[0])
	n := len(grid)
	if k/(m*n) > 0 && (m*n)%k == 0 {
		return grid
	}

	var result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < n; i++ {

		for j := 0; j < m; j++ {
			curm := j + k
			curn := i
			for curm >= m {
				curm -= m
				curn++
			}
			for curn >= n {
				curn -= n
			}

			result[curn][curm] = grid[i][j]
		}

	}

	return result
}
