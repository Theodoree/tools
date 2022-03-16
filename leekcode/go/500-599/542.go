package _00_599




/*
542. 01 矩阵
给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
两个相邻元素间的距离为 1 。
*/
func updateMatrix(mat [][]int) [][]int {
	visit:= make([][]bool,len(mat))
	for i:=0;i<len(mat);i++{
		visit[i] = make([]bool,len(mat[0]))
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] >= 1 {
				continue
			}
			bfs(mat,visit,[]int{i,j,0})

		}
	}
	return mat
}

func bfs(mat [][]int,visit [][]bool, arr []int)  {//
	var subArr []int
	for len(arr) > 0 {
		for j := 0; j < len(arr)/3; j++ {
			r := arr[j*3]
			c := arr[j*3+1]
			cnt := arr[j*3+2]

			if visit[r][c] && cnt >= mat[r][c] { // 检查过,且小于或者等于当前节点,那么就没有必要继续检查
				continue
			}
			if mat[r][c] != 0 {
				visit[r][c] = true
				mat[r][c] = cnt
			}

			// 上
			if r-1 >= 0 && mat[r-1][c] != 0{
				subArr = append(subArr, r-1, c, cnt+1)
			}

			// 下
			if r+1 < len(mat) && mat[r+1][c] != 0{
				subArr = append(subArr, r+1, c, cnt+1)
			}

			// 左
			if c-1 >= 0 && mat[r][c-1] != 0{
				subArr = append(subArr, r, c-1, cnt+1)
			}

			// 右
			if c+1 < len(mat[0]) && mat[r][c+1] != 0{
				subArr = append(subArr, r, c+1, cnt+1)
			}
		}
		arr,subArr = subArr,arr
		subArr = subArr[:0]
	}

	return
}

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
