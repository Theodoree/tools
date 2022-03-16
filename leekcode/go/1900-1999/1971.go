package _900_1999

/*
1971. 寻找图中是否存在路径
有一个具有 n个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）。图中的边用一个二维整数数组 edges 表示，其中 edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。 每个顶点对由 最多一条 边连接，并且没有顶点存在与自身相连的边。
请你确定是否存在从顶点 start 开始，到顶点 end 结束的 有效路径 。
给你数组 edges 和整数 n、start和end，如果从 start 到 end 存在 有效路径 ，则返回 true，否则返回 false 。
*/
func find(h []int, x int) int {
	if h[x] != x {
		h[x] = find(h, h[x])
	}
	return h[x]
}

func merge(h []int, x, y int) {
	x, y = find(h, x), find(h, y)
	if x == y {
		return
	}
	if x < y {
		x, y = y, x
	}
	h[x] = y
	return
}

func validPath(n int, edges [][]int, start int, end int) bool {
	h := make([]int, n+1)
	for i := 0; i < len(h); i++ {
		h[i] = i
	}
	for _, v := range edges {
		merge(h, v[0], v[1])
	}
	return find(h, start) == find(h, end)
}
