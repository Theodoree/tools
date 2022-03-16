package _500_1599

/*
1572. 矩阵对角线元素的和
给你一个正方形矩阵 mat，请你返回矩阵对角线元素的和。
请你返回在矩阵主对角线上的元素和副对角线上且不在主对角线上元素的和。
*/
func diagonalSum(mat [][]int) int {

	var sum int
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i]
		mat[i][i] = 0
		sum += mat[i][len(mat)-i-1]
	}

	return sum

}
