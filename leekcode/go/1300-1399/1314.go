package _300_1399


/*
1314. 矩阵区域和
给你一个 m x n 的矩阵 mat 和一个整数 k ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素 mat[r][c] 的和：
i - k <= r <= i + k,
j - k <= c <= j + k 且
(r, c) 在矩阵内。
*/

func matrixBlockSum(mat [][]int, k int) [][]int {

	for i:=0;i<len(mat);i++{
		for j:=1;j<len(mat[i]);j++{
			mat[i][j] += mat[i][j-1]
		}
	}

	result:=make([][]int,len(mat))
	for idx:=range result{
		result[idx] = make([]int,len(mat[0]))
	}

	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[i]);j++{

			lRow:=i-k
			rRow:=i+k
			if lRow < 0 {
				lRow = 0
			}
			if rRow >= len(mat){
				rRow  = len(mat)-1
			}



			lCol:=j-k-1
			rCol:=j+k
			if rCol >= len(mat[i]){
				rCol = len(mat[i])-1
			}

			for ;lRow<=rRow;lRow++{
				if lCol < 0 {
					result[i][j] += mat[lRow][rCol]
				}else{
					result[i][j] += mat[lRow][rCol] - mat[lRow][lCol]
				}

			}

		}
	}

	return result
}

