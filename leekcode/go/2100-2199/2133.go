package _100_2199



/*
2133. 检查是否每一行每一列都包含全部整数
对一个大小为 n x n 的矩阵而言，如果其每一行和每一列都包含从 1 到 n 的 全部 整数（含 1 和 n），则认为该矩阵是一个 有效 矩阵。
给你一个大小为 n x n 的整数矩阵 matrix ，请你判断矩阵是否为一个有效矩阵：如果是，返回 true ；否则，返回 false 。
*/

func checkValid(matrix [][]int) bool {


	var target [2]int
	for i:=1;i<=len(matrix);i++{
		if i >=64 {
			target[1] |=  1 << (i-64)
		}else{
			target[0] |=  1 << i
		}
	}


	for i:=0;i<len(matrix);i++{
		var cur [2]int

		for j:=0;j<len(matrix[i]);j++{ // 检查row
			num:=matrix[i][j]
			if num >=64 {
				cur[1] |=  1 << (num-64)
			}else{
				cur[0] |=  1 << num
			}
		}
		if cur[0] != target[0] || cur[1] != target[1] {
			return false
		}

		cur = [2]int{}
		for j:=0;j<len(matrix);j++{ // 检查row
			num:=matrix[j][i]
			if num >=64 {
				cur[1] |=  1 << (num-64)
			}else{
				cur[0] |=  1 << num
			}
		}

		if cur[0] != target[0] || cur[1] != target[1] {
			return false
		}

	}
	return true
}

