package _300_1399

import "math"

/*
1337. 矩阵中战斗力最弱的 K 行
给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。

请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。

如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。

军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。
*/
func kWeakestRows(mat [][]int, k int) []int {

	m := len(mat)
	n := len(mat[0])
	tmp := make([]int, 0, m)

	for i := 0; i < m; i++ {

		left := 0
		right := n - 1
		for left < right {
			mid := (left + right) / 2
			if mat[i][mid] == 0 {
				right--
			} else {
				left++
			}
		}

		for right > 0 && mat[i][right] == 0 {
			right--
		}
		if right >= 0 && mat[i][right] == 1 {
			tmp = append(tmp, right+1)
		} else {
			tmp = append(tmp, 0)
		}
	}

	result := make([]int, 0, k)

	for k > 0 {
		min := math.MaxInt64
		var minIdx int

		for i := 0; i < len(tmp); i++ {
			if tmp[i] == -1 {
				continue
			}
			if tmp[i] < min {
				min = tmp[i]
				minIdx = i
			}
		}
		result = append(result, minIdx)
		tmp[minIdx] = -1
		k--
	}

	return result
}
