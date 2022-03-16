package _00_699

/*
695. 岛屿的最大面积
给你一个大小为 m x n 的二进制矩阵 grid 。
岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
岛屿的面积是岛上值为 1 的单元格的数目。
计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
*/
func maxAreaOfIsland(grid [][]int) int {
	var max int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			count := bfs(grid, i, j)
			if count > max {
				max = count
			}
		}
	}
	return max
}

func bfs(grid [][]int, sr, sc int) int {

	ids := []int{sr, sc}

	var count int
	count++
	grid[sr][sc] = 0
	for len(ids) > 0 {
		sr = ids[len(ids)-2]
		sc = ids[len(ids)-1]
		ids = ids[:len(ids)-2]

		// up
		if sr-1 >= 0 && grid[sr-1][sc] == 1 {
			ids = append(ids, sr-1, sc)
			grid[sr-1][sc] = 0
			count++
		}

		// down
		if sr+1 < len(grid) && grid[sr+1][sc] == 1 {
			ids = append(ids, sr+1, sc)
			grid[sr+1][sc] = 0
			count++
		}

		// left
		if sc-1 >= 0 && grid[sr][sc-1] == 1 {
			ids = append(ids, sr, sc-1)
			grid[sr][sc-1] = 0
			count++
		}

		// right
		if sc+1 < len(grid[sr]) && grid[sr][sc+1] == 1 {
			ids = append(ids, sr, sc+1)
			grid[sr][sc+1] = 0
			count++
		}

	}
	return count

}
