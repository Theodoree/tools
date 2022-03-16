package _00_799

func countCornerRectangles(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	count := make([][]int, len(grid[0]))
	for i := 0; i < len(count); i++ {
		count[i] = make([]int, len(grid[0]))
	}

	var ans int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				for k := j+1; k < len(grid[i]); k++ {
					if grid[i][k] == 1 {
						ans += count[j][k]
						count[j][k]++
					}
				}
			}
		}
	}

	return ans

}


