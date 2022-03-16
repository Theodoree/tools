package _200_1299


/*
1254. 统计封闭岛屿的数目
二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
请返回 封闭岛屿 的数目。
*/
func closedIsland(grid [][]int) int {


	var cnt int // 用两个int表示也可以压缩这些状态,可以没必要
	var flag = make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		flag[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if flag[i][j] || grid[i][j] == 1 {
				continue
			}
			if bfs(flag, grid, i, j) {
				cnt++
			}
		}
	}
	return cnt
}

func bfs(flag [][]bool, grid1 [][]int, row, col int) bool {

	ok := true
	flag[row][col] = true
	idx := []int{row, col}
	var tmp []int
	for len(idx) > 0 {

		for i := 0; i < len(idx)/2; i++ {
			r := idx[i*2]
			c := idx[i*2+1]
			if grid1[r][c] == 1 {
				continue
			}
			// up
			if r-1 >= 0 && !flag[r-1][c] {
				flag[r-1][c] = true
				tmp = append(tmp, r-1, c)
			}

			// down
			if r+1 < len(grid1) && !flag[r+1][c] {
				flag[r+1][c] = true
				tmp = append(tmp, r+1, c)
			}

			// left
			if c-1 >= 0 && !flag[r][c-1] {
				flag[r][c-1] = true
				tmp = append(tmp, r, c-1)
			}

			// right
			if c+1 < len(grid1[0]) && !flag[r][c+1] {
				flag[r][c+1] = true
				tmp = append(tmp, r, c+1)
			}
			if r-1 < 0 || c-1<0 || r+1>=len(grid1) || c+1 >=len(grid1[0]){
				ok = false
			}
		}
		idx,tmp = tmp,idx
		tmp = tmp[:0]
	}

	return ok
}
