package _400_1499

/*
1496. 判断路径是否相交
给你一个字符串 path，其中 path[i] 的值可以是 'N'、'S'、'E' 或者 'W'，分别表示向北、向南、向东、向西移动一个单位。
机器人从二维平面上的原点 (0, 0) 处开始出发，按 path 所指示的路径行走。
如果路径在任何位置上出现相交的情况，也就是走到之前已经走过的位置，请返回 True ；否则，返回 False 。
*/
func isPathCrossing(path string) bool {
	var m = make(map[string]struct{})

	var x, y int
	m["0_0"] = struct{}{}

	for _, v := range path {
		switch v {
		case 'N': // NOUTH 北
			y++
		case 'S': // SOUTH 南
			y--
		case 'E': // EAST 东
			x--
		case 'W': // WEST 西
			x++
		}
		str := fmt.Sprintf("%d_%d", x, y)
		if _, ok := m[str]; ok {
			return true
		}
		m[str] = struct{}{}
	}
	return false
}
