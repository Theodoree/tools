package _00_399

import "sort"

/*
378. 有序矩阵中第 K 小的元素
给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。
*/
func kthSmallest1(matrix [][]int, k int) int {
	var row int
	var col int
	var l = make([]int, len(matrix))
	for k > 0 {

		var r int
		val := math.MaxInt64
		var c int
		for i := 0; i < len(l); i++ {
			if l[i] == -1 {
				continue
			}
			if matrix[i][l[i]] < val {
				val = matrix[i][l[i]]
				r = i
				c = l[i]
			}
		}

		row = r
		col = c
		l[r]++
		if l[r] == len(matrix[0]) {
			l[r] = -1
		}

		k--
	}

	return matrix[row][col]

}

func kthSmallest(matrix [][]int, k int) int {
	f := make([]int, 0, len(matrix)*len(matrix[0]))
	for _, v := range matrix {
		f = append(f, v...)
	}
	sort.Ints(f)
	return f[k-1]
}
