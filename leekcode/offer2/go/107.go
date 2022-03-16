package _go

func updateMatrix(mat [][]int) [][]int { // dp的思路的确好O(n)^2
	row, col := len(mat), len(mat[0])
	ret := make([][]int, row)
	for i := range ret {
		ret[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if mat[i][j] != 0 {
				ret[i][j] = math.MaxInt32
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i >= 1 {
				ret[i][j] = min(ret[i][j], ret[i-1][j]+1)
			}
			if j >= 1 {
				ret[i][j] = min(ret[i][j], ret[i][j-1]+1)
			}
		}
	}

	for i := row - 1; i >= 0; i-- {
		for j := 0; j < col; j++ {
			if i < row-1 {
				ret[i][j] = min(ret[i][j], ret[i+1][j]+1)
			}
			if j >= 1 {
				ret[i][j] = min(ret[i][j], ret[i][j-1]+1)
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := col - 1; j >= 0; j-- {
			if i >= 1 {
				ret[i][j] = min(ret[i][j], ret[i-1][j]+1)
			}
			if j < col-1 {
				ret[i][j] = min(ret[i][j], ret[i][j+1]+1)
			}
		}
	}

	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if i < row-1 {
				ret[i][j] = min(ret[i][j], ret[i+1][j]+1)
			}
			if j < col-1 {
				ret[i][j] = min(ret[i][j], ret[i][j+1]+1)
			}
		}
	}

	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
