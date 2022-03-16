package _go

/*
剑指 Offer 12. 矩阵中的路径

请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用下画线标出）。

[["a","b","c","e"],
["s","f","c,"s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。



示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false
提示：

1 <= board.length <= 200
1 <= board[i].length <= 200
*/

func exist(board [][]byte, word string) bool {

	buf := []byte(word)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == buf[0] {
				if dfs(buf, 0, i, j, board) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(word []byte, index, i, j int, board [][]byte) bool {
	if index >= len(word) {
		return true
	}

	if i >= len(board) || j >= len(board[0]) || i < 0 || j < 0 || board[i][j] != word[index] {
		return false
	}

	board[i][j] += 100
	result := dfs(word, index+1, i+1, j, board) || dfs(word, index+1, i-1, j, board) || dfs(word, index+1, i, j+1, board) || dfs(word, index+1, i, j-1, board)
	board[i][j] -= 100

	return result
}
