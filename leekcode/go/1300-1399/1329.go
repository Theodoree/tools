package _300_1399



/*
1329. 将矩阵按对角线排序
矩阵对角线 是一条从矩阵最上面行或者最左侧列中的某个元素开始的对角线，沿右下方向一直到矩阵末尾的元素。例如，矩阵 mat 有 6 行 3 列，从 mat[2][0] 开始的 矩阵对角线 将会经过 mat[2][0]、mat[3][1] 和 mat[4][2] 。
给你一个 m * n 的整数矩阵 mat ，请你将同一条 矩阵对角线 上的元素按升序排序后，返回排好序的矩阵。
*/

func diagonalSort(mat [][]int) [][]int {

	startRow := len(mat) - 1
	startCol := 0
	arr:=make([]int,0,len(mat[0]))
	for startCol < len(mat[0]){
		arr = arr[:0]
		r:=startRow
		c:=startCol
		for r < len(mat) && c < len(mat[0]) {
			arr = append(arr,mat[r][c])
			r++
			c++
		}
		sort.Ints(arr)

		i:=0
		r=startRow
		c=startCol
		for r < len(mat) && c < len(mat[0]) {
			mat[r][c] = arr[i]
			r++
			c++
			i++
		}

		if startRow == 0 {
			startCol++
		} else {
			startRow--
		}

	}
	return mat
}

