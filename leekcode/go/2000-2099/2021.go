package _000_2099


/*
2021. 街上最亮的位置
一条街上有很多的路灯，路灯的坐标由数组 lights 的形式给出。 每个 lights[i] = [positioni, rangei] 代表坐标为 positioni 的路灯照亮的范围为 [positioni - rangei, positioni + rangei] （包括顶点）。
位置 p 的亮度由能够照到 p的路灯的数量来决定的。
给出 lights, 返回最亮的位置 。如果有很多，返回坐标最小的。
*/
func brightestPosition(lights [][]int) (pos int) {
	a := make([]int, 0, len(lights)*2)
	for _, l := range lights {
		p, d := l[0], l[1]
		a = append(a, (p-d)<<1|1, (p+d+1)<<1) // 最低位存储是区间左端点还是区间右端点
	}
	sort.Ints(a)

	s, maxS := 0, 0
	for i, v := range a {
		s += v&1<<1 - 1 // 根据最低位来 +1 或 -1
		if (i == len(a)-1 || a[i+1]>>1 != v>>1) && s > maxS {
			maxS = s
			pos = v >> 1
		}
	}
	return
}

