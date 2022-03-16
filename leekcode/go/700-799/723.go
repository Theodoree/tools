package _00_799



/*
723. 粉碎糖果
这个问题是实现一个简单的消除算法。
给定一个 m x n 的二维整数数组 board 代表糖果所在的方格，不同的正整数 board[i][j] 代表不同种类的糖果，如果 board[i][j] == 0 代表 (i, j) 这个位置是空的。
给定的方格是玩家移动后的游戏状态，现在需要你根据以下规则粉碎糖果，使得整个方格处于稳定状态并最终输出：
如果有三个及以上水平或者垂直相连的同种糖果，同一时间将它们粉碎，即将这些位置变成空的。
在同时粉碎掉这些糖果之后，如果有一个空的位置上方还有糖果，那么上方的糖果就会下落直到碰到下方的糖果或者底部，这些糖果都是同时下落，也不会有新的糖果从顶部出现并落下来。
通过前两步的操作，可能又会出现可以粉碎的糖果，请继续重复前面的操作。
当不存在可以粉碎的糖果，也就是状态稳定之后，请输出最终的状态。
你需要模拟上述规则并使整个方格达到稳定状态，并输出。
*/
func candyCrush(board [][]int) [][]int {
	rows := len(board)
	cols := len(board[0])

	var todo bool
	//一行一行的扫描
	for r := 0; r < rows; r++ {
		for c := 0; c+2 < cols; c++ {
			v := abs(board[r][c])
			if v != 0 && v == abs(board[r][c+1]) && v == abs(board[r][c+2]) {
				board[r][c], board[r][c+1], board[r][c+2] = -v, -v, -v
				todo = true
			}
		}
	}

	//一列一列的扫描

	for r := 0; r+2 < rows; r++ {
		for c := 0; c < cols; c++ {
			v := abs(board[r][c])
			if v != 0 && v == abs(board[r+1][c]) && v == abs(board[r+2][c]) {
				board[r][c], board[r+1][c], board[r+2][c] = -v, -v, -v
				todo = true
			}
		}
	}

	//粉碎
	if todo {
		for c := 0; c < cols; c++ {
			wr := rows - 1
			for r := rows - 1; r >= 0; r-- {
				if board[r][c] > 0 {
					board[wr][c] = board[r][c]
					wr--
				}
			}
			for wr >= 0 {
				board[wr][c] = 0
				wr--
			}
		}
	}

	if todo {
		candyCrush(board)
	}
	return board
}


func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

