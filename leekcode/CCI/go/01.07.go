package _go

/*
面试题 01.07. 旋转矩阵
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
不占用额外内存空间能否做到？
*/
func rotate(matrix [][]int) {
	high := len(matrix)
	width := high
	for i := 0; i < high/2; i++ {
		for j := 0; j < width; j++ {
			matrix[i][j], matrix[high-i-1][j] = matrix[high-i-1][j], matrix[i][j]
		}
	}
	for i := 0; i < high; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
