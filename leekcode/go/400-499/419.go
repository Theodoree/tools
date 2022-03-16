package _00_499

/*
419. 甲板上的战舰
给你一个大小为 m x n 的矩阵 board 表示甲板，其中，每个单元格可以是一艘战舰 'X' 或者是一个空位 '.' ，返回在甲板 board 上放置的 战舰 的数量。
战舰 只能水平或者垂直放置在 board 上。换句话说，战舰只能按 1 x k（1 行，k 列）或 k x 1（k 行，1 列）的形状建造，其中 k 可以是任意大小。两艘战舰之间至少有一个水平或垂直的空位分隔 （即没有相邻的战舰）。
*/

func countBattleships(board [][]byte) int {
	var cnt int
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[i]);j++{
			if board[i][j] == '.' || board[i][j] == 'Z'{
				continue
			}
			board[i][j] = 'Z'

			if i-1 >=0 && board[i-1][j] == 'Z' {
				continue
			}
			if i+1 <len(board)&& board[i+1][j] == 'Z' {
				continue
			}

			if j-1 >=0 && board[i][j-1] == 'Z' {
				continue
			}
			if j+1 <len(board[i])&& board[i][j+1] == 'Z' {
				continue
			}
			cnt++
		}
	}
	return cnt
}

