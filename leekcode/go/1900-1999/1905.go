package _900_1999



/*
1905. 统计子岛屿
给你两个 m x n 的二进制矩阵 grid1 和 grid2 ，它们只包含 0 （表示水域）和 1 （表示陆地）。一个 岛屿 是由 四个方向 （水平或者竖直）上相邻的 1 组成的区域。任何矩阵以外的区域都视为水域。

如果 grid2 的一个岛屿，被 grid1 的一个岛屿 完全 包含，也就是说 grid2 中该岛屿的每一个格子都被 grid1 中同一个岛屿完全包含，那么我们称 grid2 中的这个岛屿为 子岛屿 。

请你返回 grid2 中 子岛屿 的 数目 。
*/
func countSubIslands(grid1 [][]int, grid2 [][]int) int {

	grid1, grid2 = grid2, grid1

	var cnt int
	var flag = make([][]bool, len(grid1))

	for i := 0; i < len(grid1); i++ {
		flag[i] = make([]bool, len(grid1[0]))
	}

	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			if flag[i][j] || grid1[i][j] == 0 {
				continue
			}
			if bfs(flag, grid1, grid2, i, j) {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
	return cnt
}

func bfs(flag [][]bool, grid1, grid2 [][]int, row, col int) bool {

	// grid1和grid2同时都为1
	ok := true
	flag[row][col] = true
	idx := []int{row, col}
	var tmp []int
	for len(idx) > 0 {

		for i := 0; i < len(idx)/2; i++ {
			r := idx[i*2]
			c := idx[i*2+1]
			if grid1[r][c] == 1 && grid2[r][c] == 0 {
				ok = false
			}
			// up
			if r-1 >= 0 && !flag[r-1][c] && grid1[r-1][c] == 1 {
				flag[r-1][c] = true
				tmp = append(tmp, r-1, c)
			}

			// down
			if r+1 < len(grid1) && !flag[r+1][c] && grid1[r+1][c] == 1 {
				flag[r+1][c] = true
				tmp = append(tmp, r+1, c)
			}

			// left
			if c-1 >= 0 && !flag[r][c-1] && grid1[r][c-1] == 1 {
				flag[r][c-1] = true
				tmp = append(tmp, r, c-1)
			}

			// right
			if c+1 < len(grid1[0]) && !flag[r][c+1] && grid1[r][c+1] == 1 {
				flag[r][c+1] = true
				tmp = append(tmp, r, c+1)
			}
		}
		idx,tmp = tmp,idx
		tmp = tmp[:0]
	}

	return ok
}
