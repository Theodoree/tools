package _200_1299

/*
1252. 奇数值单元格的数目
给你一个 m x n 的矩阵，最开始的时候，每个单元格中的值都是 0。
另有一个二维索引数组 indices，indices[i] = [ri, ci] 指向矩阵中的某个位置，其中 ri 和 ci 分别表示指定的行和列（从 0 开始编号）。
对 indices[i] 所指向的每个位置，应同时执行下述增量操作：
ri 行上的所有单元格，加 1 。
ci 列上的所有单元格，加 1 。
给你 m、n 和 indices 。请你在执行完所有 indices 指定的增量操作后，返回矩阵中 奇数值单元格 的数目。
*/
func oddCells(m int, n int, indices [][]int) int {
	var buf = make([][]int, m)
	for idx := range buf {
		buf[idx] = make([]int, n)
	}

	for i := 0; i < len(indices); i++ {
		r := indices[i][0]
		c := indices[i][1]
		for j := 0; j < len(buf[r]); j++ {
			buf[r][j]++
		}
		for j := 0; j < len(buf); j++ {
			buf[j][c]++
		}
	}

	var count int
	for i := 0; i < len(buf); i++ {
		for _, v := range buf[i] {
			if v%2 != 0 {
				count++
			}
		}
	}
	return count

}
