package __99


/*
79. 单词搜索


给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true.
给定 word = "SEE", 返回 true.
给定 word = "ABCB", 返回 false.
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

    if i >= len(board) || j >= len(board[0]) ||  i < 0 || j < 0 || board[i][j] != word[index] {
        return false
    }

    board[i][j] += 100
    result := dfs(word, index+1, i+1, j, board) || dfs(word, index+1, i-1, j, board) || dfs(word, index+1, i, j+1, board) || dfs(word, index+1, i, j-1, board)
    board[i][j] -= 100

    return result
}
