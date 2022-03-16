package _000_2099

/*
2022. 将一维数组转变成二维数组
给你一个下标从 0 开始的一维整数数组 original 和两个整数 m 和  n 。你需要使用 original 中 所有 元素创建一个 m 行 n 列的二维数组。
original 中下标从 0 到 n - 1 （都 包含 ）的元素构成二维数组的第一行，下标从 n 到 2 * n - 1 （都 包含 ）的元素构成二维数组的第二行，依此类推。
请你根据上述过程返回一个 m x n 的二维数组。如果无法构成这样的二维数组，请你返回一个空的二维数组。
*/
func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	if m == 1 {
		return [][]int{original}
	}

	var result [][]int
	for i := 0; i < m; i++ {
		result = append(result, make([]int, 0, n))
	}

	for i := 0; i < len(original); i++ {
		result[i/n] = append(result[i/n], original[i])
	}
	return result
}
