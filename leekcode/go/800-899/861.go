package _00_899



/*
861. 翻转矩阵后的得分
有一个二维矩阵 A 其中每个元素的值为 0 或 1 。
移动是指选择任一行或列，并转换该行或列中的每一个值：将所有 0 都更改为 1，将所有 1 都更改为 0。
在做出任意次数的移动后，将该矩阵的每一行都按照二进制数来解释，矩阵的得分就是这些数字的总和。
返回尽可能高的分数。
*/
func matrixScore(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}


	row:=len(grid)
	col:=len(grid[0])


	// 先检查是否需要翻转
	for i:=0;i<row;i++{
		if grid[i][0] == 0 { // 翻转
			for j:=0;j<col;j++{
				if grid[i][j] == 0 {
					grid[i][j] = 1
				}else {
					grid[i][j] = 0
				}
			}
		}
	}



	// 检查列是否需要反转
	for i:=1;i<col;i++{
		var old int
		var reverse int
		for j:=0;j<row;j++{
			if grid[j][i] == 1 {
				old+=  1 << col-i
			}else {
				reverse += 1 << col-i
			}
		}
		if reverse > old {
			for j:=0;j<len(grid);j++{
				if grid[j][i] == 1 {
					grid[j][i] = 0
				}else {
					grid[j][i] = 1
				}
			}
		}
	}


	var sum int

	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			if  grid[i][j] == 1 {
				sum+= 1 << (col-j-1)
			}

		}

	}

	return sum
}

