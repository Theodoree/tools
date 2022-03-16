package _00_399

/*
311. 稀疏矩阵的乘法
给定两个 稀疏矩阵 ：大小为 m x k 的稀疏矩阵 mat1 和大小为 k x n 的稀疏矩阵 mat2 ，返回 mat1 x mat2 的结果。你可以假设乘法总是可能的。
示例 1：
输入：mat1 = [[1,0,0],[-1,0,3]], mat2 = [[7,0,0],[0,0,0],[0,0,1]]
输出：[[7,0,0],[-7,0,3]]
 示例 2:
输入：mat1 = [[0]], mat2 = [[0]]
输出：[[0]]
*/
func multiply(A [][]int, B [][]int) [][]int {
	ha, wa, wb := len(A), len(A[0]), len(B[0])
	out := make([][]int, ha)
	for i := 0; i < ha; i++ {
		out[i] = make([]int, wb)
	}
	for i := 0; i < ha; i++ {
		for j := 0; j < wa; j++ {
			a := A[i][j]
			if a == 0 {
				continue
			}
			br := B[j]
			for k := 0; k < wb; k++ {
				if br[k] == 0 {
					continue
				}
				out[i][k] += a * br[k]
			}
		}
	}
	return out
}

