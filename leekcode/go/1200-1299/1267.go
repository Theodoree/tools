package _200_1299

/*
1267. 统计参与通信的服务器
这里有一幅服务器分布图，服务器的位置标识在 m * n 的整数矩阵网格 grid 中，1 表示单元格上有服务器，0 表示没有。

如果两台服务器位于同一行或者同一列，我们就认为它们之间可以进行通信。

请你统计并返回能够与至少一台其他服务器进行通信的服务器的数量。
*/
func countServers(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	var row = make([]int, len(grid))
	var col = make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}

	var cnt int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && (row[i]|col[j] >= 2) {
				cnt++
			}
		}
	}

	return cnt
}
