package _go

/*
剑指 Offer II 105. 岛屿的最大面积
给定一个由 0 和 1 组成的非空二维数组 grid ，用来表示海洋岛屿地图。
一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
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
