package _00_999


/*
994. 腐烂的橘子
在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
*/

func orangesRotting(grid [][]int) int {

	var result []int
	var cnt int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				result = append(result, i, j)
			}
			cnt += grid[i][j] & 1
		}
	}

	var count int
	var tmp []int
	for len(result) > 0 && cnt > 0 {

		for i := 0; i < len(result)/2; i++ {
			row := result[i*2]
			col := result[i*2+1]

			// 上
			if row-1 >= 0 && grid[row-1][col] == 1 {
				cnt--
				grid[row-1][col] = 2
				tmp = append(tmp, row-1, col)
			}

			// 下
			if row+1 < len(grid) && grid[row+1][col] == 1 {
				cnt--
				grid[row+1][col] = 2
				tmp = append(tmp, row+1, col)
			}

			// 左
			if col-1 >= 0 && grid[row][col-1] == 1 {
				cnt--
				grid[row][col-1] = 2
				tmp = append(tmp, row, col-1)
			}

			// 右
			if col+1 < len(grid[0]) && grid[row][col+1] == 1 {
				cnt--
				grid[row][col+1] = 2
				tmp = append(tmp, row, col+1)
			}

		}

		tmp, result = result, tmp
		tmp = tmp[:0]

		count++
	}
	if cnt > 0 {
		return -1
	}
	return count

}
