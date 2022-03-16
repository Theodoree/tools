package _100_1199

/*
1184. 公交站间的距离
环形公交路线上有 n 个站，按次序从 0 到 n - 1 进行编号。我们已知每一对相邻公交站之间的距离，distance[i] 表示编号为 i 的车站和编号为 (i + 1) % n 的车站之间的距离。

环线上的公交车都可以按顺时针和逆时针的方向行驶。

返回乘客从出发点 start 到目的地 destination 之间的最短距离。
*/
func distanceBetweenBusStops(distance []int, start int, destination int) int {

	// 顺时针
	var sum, sum1 int

	cur := start
	for cur%len(distance) != destination {
		sum += distance[cur%len(distance)]
		cur++
	}

	// 逆时针
	cur = start + len(distance)
	for cur%len(distance) != destination {
		cur--
		sum1 += distance[cur%len(distance)]
	}

	if sum1 > sum {
		return sum
	}

	return sum1
}
