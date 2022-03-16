package _00_399


/*
348. 设计井字棋
请在 n × n 的棋盘上，实现一个判定井字棋（Tic-Tac-Toe）胜负的神器，判断每一次玩家落子后，是否有胜出的玩家。
在这个井字棋游戏中，会有 2 名玩家，他们将轮流在棋盘上放置自己的棋子。
在实现这个判定器的过程中，你可以假设以下这些规则一定成立：
      1. 每一步棋都是在棋盘内的，并且只能被放置在一个空的格子里；
      2. 一旦游戏中有一名玩家胜出的话，游戏将不能再继续；
      3. 一个玩家如果在同一行、同一列或者同一斜对角线上都放置了自己的棋子，那么他便获得胜利。
*/

const (
	Col = 0
	Row = 1
	Angle = 2
)

type TicTacToe struct {
	event [3][3][]int // [playid][event]
	n int

}


func Constructor(n int) TicTacToe {
	var t TicTacToe
	for i:=range t.event{
		for j:=range t.event[i]{
			if j == Angle {
				t.event[i][j] = make([]int,2)
				continue
			}
			t.event[i][j] = make([]int,n)
		}
	}
	t.n = n
	return t
}


func (this *TicTacToe) Move(row int, col int, player int) int {
	this.event[player][Row][row]++
	this.event[player][Col][col]++
	if row == col {
		this.event[player][Angle][0]++
	}
	if row + col == this.n-1 {
		this.event[player][Angle][1]++
	}

	if  this.event[player][Row][row] == this.n  || this.event[player][Col][col] == this.n {
		return player
	}
	if  this.event[player][Angle][0] == this.n  || this.event[player][Angle][1] == this.n {
		return player
	}


	return 0
}
