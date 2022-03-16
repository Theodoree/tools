package _200_1299

/*
1275. 找出井字棋的获胜者
A 和 B 在一个 3 x 3 的网格上玩井字棋。
井字棋游戏的规则如下：
玩家轮流将棋子放在空方格 (" ") 上。
第一个玩家 A 总是用 "X" 作为棋子，而第二个玩家 B 总是用 "O" 作为棋子。
"X" 和 "O" 只能放在空方格中，而不能放在已经被占用的方格上。
只要有 3 个相同的（非空）棋子排成一条直线（行、列、对角线）时，游戏结束。
如果所有方块都放满棋子（不为空），游戏也会结束。
游戏结束后，棋子无法再进行任何移动。
给你一个数组 moves，其中每个元素是大小为 2 的另一个数组（元素分别对应网格的行和列），它按照 A 和 B 的行动顺序（先 A 后 B）记录了两人各自的棋子位置。
如果游戏存在获胜者（A 或 B），就返回该游戏的获胜者；如果游戏以平局结束，则返回 "Draw"；如果仍会有行动（游戏未结束），则返回 "Pending"。
你可以假设 moves 都 有效（遵循井字棋规则），网格最初是空的，A 将先行动。
*/
func tictactoe(moves [][]int) string {
	var buf [3][3]byte

	for i := 0; i < len(moves); i++ {
		if i%2 == 0 {
			buf[moves[i][0]][moves[i][1]] = 'A'
		} else {
			buf[moves[i][0]][moves[i][1]] = 'B'
		}
	}

	for i := 0; i < 3; i++ {

		if i == 0 && buf[0][0] > 0 && buf[0][0] == buf[1][1] && buf[0][0] == buf[2][2] {
			return string(buf[0][0])
		}

		if i == 2 && buf[2][0] > 0 && buf[2][0] == buf[1][1] && buf[1][1] == buf[0][2] {
			return string(buf[2][0])
		}

		if buf[i][0] > 0 && buf[i][0] == buf[i][1] && buf[i][0] == buf[i][2] {
			return string(buf[i][0])
		}

		if buf[0][i] > 0 && buf[0][i] == buf[1][i] && buf[0][i] == buf[2][i] {
			return string(buf[0][i])
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}
