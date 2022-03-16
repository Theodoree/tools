package _00_299

/*
289. 生命游戏
根据 百度百科 ， 生命游戏 ，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。
给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：
如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。给你 m x n 网格面板 board 的当前状态，返回下一个状态。
*/

func check(board [][]int, r, c int, n int) int {
	var cnt int
	// 上
	if r-1 >= 0 && board[r-1][c] == n {
		cnt++
	}
	// 下
	if r+1 < len(board) && board[r+1][c] == n {
		cnt++
	}

	// 左
	if c-1 >= 0 && board[r][c-1] == n {
		cnt++
	}
	//右
	if c+1 < len(board[0]) && board[r][c+1] == n {
		cnt++
	}

	// 左上
	if c-1 >= 0 && r-1 >= 0 && board[r-1][c-1] == n {
		cnt++
	}
	// 右下
	if r+1 < len(board) && c+1 < len(board[0]) && board[r+1][c+1] == n {
		cnt++
	}

	// 左下
	if r+1 < len(board) && c-1 >= 0 && board[r+1][c-1] == n {
		cnt++
	}

	// 右上
	if c+1 < len(board[0]) && r-1 >= 0 && board[r-1][c+1] == n {
		cnt++
	}

	return cnt
}

func gameOfLife(board [][]int) {
	b := make([][]int, len(board))
	for idx := range b {
		b[idx] = make([]int, len(board[0]))
		copy(b[idx], board[idx])
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			switch board[i][j] {
			case 0: // 检查是否可以复活
				cnt := check(b, i, j, 1)
				if cnt == 3 {
					board[i][j] = 1
				}

			case 1: // 检查是否会死
				cnt := check(b, i, j, 1)
				if !(cnt == 2 || cnt == 3) {
					board[i][j] = 0
				}

			}
		}
	}
}
