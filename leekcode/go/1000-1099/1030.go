package _000_1099

/*
1030. 距离顺序排列矩阵单元格
给定四个整数 row ,   cols ,  rCenter 和 cCenter 。有一个 rows x cols 的矩阵，你在单元格上的坐标是 (rCenter, cCenter) 。
返回矩阵中的所有单元格的坐标，并按与 (rCenter, cCenter) 的 距离 从最小到最大的顺序排。你可以按 任何 满足此条件的顺序返回答案。
单元格(r1, c1) 和 (r2, c2) 之间的距离为|r1 - r2| + |c1 - c2|。
*/

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	var result [][][]int = make([][][]int,rows*cols)
	var flags = make([][]bool,rows)
	for i:=0;i<len(flags);i++{
		flags[i] = make([]bool,cols)
	}
	flags[rCenter][cCenter] = true

	var idxs = []int{0,rCenter,cCenter}

	var curIdxs []int
	for  len(idxs) > 0 {

		curIdxs,idxs = idxs,curIdxs
		idxs = idxs[:0]

		for i:=0;i<len(curIdxs)/3;i++{
			resultIdx := curIdxs[i*3]
			r := curIdxs[i*3+1]
			c := curIdxs[i*3+2]

			result[resultIdx] = append(result[resultIdx], []int{r, c})

			resultIdx++
			// up
			if r-1 >= 0 && !flags[r-1][c] {
				flags[r-1][c] = true
				idxs = append(idxs, resultIdx, r-1, c)
			}

			// down
			if r+1 < rows && !flags[r+1][c] {
				flags[r+1][c] = true
				idxs = append(idxs, resultIdx, r+1, c)
			}

			// left
			if c-1 >= 0 && !flags[r][c-1] {
				flags[r][c-1] = true
				idxs = append(idxs, resultIdx, r, c-1)
			}

			// right
			if c+1 < cols && !flags[r][c+1] {
				flags[r][c+1] = true
				idxs = append(idxs, resultIdx, r, c+1)
			}
		}

	}

	var r [][]int
	for _,v:=range result{
		r = append(r,v...)
	}

	return r
}

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	maxDis := max(rCenter, rows-1-rCenter) + max(cCenter, cols-1-cCenter)
	result := make([][]int, rows*cols)
	var colDis, index int
	for dis := 0; dis <= maxDis; dis++ {
		for row := max(rCenter-dis, 0); row <= min(rCenter+dis, rows-1); row++ {
			colDis = dis - abs(row - rCenter)
			if cCenter-colDis >= 0 {
				result[index] = []int{row, cCenter - colDis}
				index++
			}
			if cCenter+colDis <= cols-1 && colDis != 0 {
				result[index] = []int{row, cCenter + colDis}
				index++
			}
		}
	}
	return result
}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


