package _00_899

/*
807. 保持城市天际线
给你一座由 n x n 个街区组成的城市，每个街区都包含一座立方体建筑。给你一个下标从 0 开始的 n x n 整数矩阵 grid ，其中 grid[r][c] 表示坐落于 r 行 c 列的建筑物的 高度 。
城市的 天际线 是从远处观察城市时，所有建筑物形成的外部轮廓。从东、南、西、北四个主要方向观测到的 天际线 可能不同。
我们被允许为 任意数量的建筑物 的高度增加 任意增量（不同建筑物的增量可能不同） 。 高度为 0 的建筑物的高度也可以增加。然而，增加的建筑物高度 不能影响 从任何主要方向观察城市得到的 天际线 。
在 不改变 从任何主要方向观测到的城市 天际线 的前提下，返回建筑物可以增加的 最大高度增量总和 。
*/
func maxIncreaseKeepingSkyline(grid [][]int) int {

	var col = make([]int, len(grid[0]))
	var row = make([]int, len(grid))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > row[i] {
				row[i] = grid[i][j]
			}
			if grid[j][i] > col[i] {
				col[i] = grid[j][i]
			}
		}

	}

	min := func(i, j int) int {
		if j > i {
			return i
		}
		return j
	}

	// 行列取最小值
	var sum int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			max := min(col[j], row[i])
			if grid[i][j] < max {
				sum += max - grid[i][j]
			}
		}
	}

	return sum
}
