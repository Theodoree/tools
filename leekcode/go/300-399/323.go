package _00_399



/*
323. 无向图中连通分量的数目
你有一个包含 n 个节点的图。给定一个整数 n 和一个数组 edges ，其中 edges[i] = [ai, bi] 表示图中 ai 和 bi 之间有一条边。
返回 图中已连接分量的数目 。
*/
func countComponents(n int, edges [][]int) int {
	ufs := newUfs(n)
	for _, v := range edges {
		ufs.Union(v[0], v[1])
		ufs.Union(v[1], v[0])
	}

	var arr = make([]int, n)
	for _, v := range ufs.roots {
		arr[v]++
	}
	return ufs.Count()
}
