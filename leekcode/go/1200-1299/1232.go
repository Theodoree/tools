package _200_1299

/*
1232. 缀点成线
在一个 XY 坐标系中有一些点，我们用数组 coordinates 来分别记录它们的坐标，其中 coordinates[i] = [x, y] 表示横坐标为 x、纵坐标为 y 的点。

请你来判断，这些点是否在该坐标系中属于同一条直线上，是则返回 true，否则请返回 false。
*/
func checkStraightLine(coordinates [][]int) bool {
	deltaX, deltaY := coordinates[0][0], coordinates[0][1]
	for _, p := range coordinates {
		p[0] -= deltaX
		p[1] -= deltaY
	}
	A, B := coordinates[1][1], -coordinates[1][0]
	for _, p := range coordinates[2:] {
		x, y := p[0], p[1]
		if A*x+B*y != 0 {
			return false
		}
	}
	return true
}
