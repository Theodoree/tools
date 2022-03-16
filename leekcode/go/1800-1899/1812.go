package _800_1899

/*
1812. 判断国际象棋棋盘中一个格子的颜色
给你一个坐标 coordinates ，它是一个字符串，表示国际象棋棋盘中一个格子的坐标。下图是国际象棋棋盘示意图。
*/
func squareIsWhite(coordinates string) bool {

	var i, i1 byte
	i = coordinates[0] - 'a' + 1
	i1 = coordinates[1] - '0'

	if i%2 == 0 {
		return i1%2 == 1
	}
	return i1%2 == 0
}
