package _00_299

/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
*/
func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				bfs(grid, [][2]int{{i, j}})
				count++
			}
		}
	}
	return count
}

func bfs(grid [][]byte, arr [][2]int) {

	for len(arr) > 0 {
		idx := arr[len(arr)-1]
		arr = arr[:len(arr)-1]

		n := idx[0]
		m := idx[1]

		grid[n][m] = '0'

		// up
		if n-1 >= 0 && grid[n-1][m] == '1' {
			arr = append(arr, [2]int{n - 1, m})
		}

		// down
		if n+1 < len(grid) && grid[n+1][m] == '1' {
			arr = append(arr, [2]int{n + 1, m})
		}

		// left
		if m-1 >= 0 && grid[n][m-1] == '1' {
			arr = append(arr, [2]int{n, m - 1})
		}
		// right
		if m+1 < len(grid[n]) && grid[n][m+1] == '1' {
			arr = append(arr, [2]int{n, m + 1})
		}
	}

}
