package _00_999




/*
999. 车的可用捕获量

在一个 8 x 8 的棋盘上，有一个白色车（rook）。也可能有空方块，白色的象（bishop）和黑色的卒（pawn）。
它们分别以字符 “R”，“.”，“B” 和 “p” 给出。大写字符表示白棋，小写字符表示黑棋。

车按国际象棋中的规则移动：它选择四个基本方向中的一个（北，东，西和南），
然后朝那个方向移动，直到它选择停止、到达棋盘的边缘或移动到同一方格来捕获该方格上颜色相反的卒。另外，车不能与其他友方（白色）象进入同一个方格。

返回车能够在一次移动中捕获到的卒的数量。

*/

func numRookCaptures(board [][]byte) int {

    for i := 0; i < len(board); i++ {
        var p int
        for j := 0; j < len(board[0]); j++ {

            switch {
            case board[i][j] == 'p':
                p = 1
            case board[i][j] == 'B':
                p = 0
            }

            if board[i][j] == 'R' {
                index1 := j + 1

                for index1 < len(board[0]) {
                    if board[i][index1] == 'B' {
                        break
                    }
                    if board[i][index1] == 'p' {
                        p++
                        break
                    }
                    index1++
                }

                index2 := i - 1
                index3 := i + 1
                for {
                    if index2 < 0 && index3 > len(board) {
                        break
                    }

                    if index2 >= 0 {
                        switch  {
                        case board[index2][j] == 'B':
                            index2 = -1
                        case board[index2][j] == 'p':
                            p++
                            index2 = - 1
                        }
                    }
                    if index3 < len(board) {
                        switch {
                        case board[index3][j] == 'B':
                            index3 = len(board)
                        case board[index3][j] == 'p':
                            p++
                            index3 = len(board)
                        }
                    }

                    index2--
                    index3++
                }

                return p

            }
        }

    }
    return 0
}
