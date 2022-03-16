package _00_299


/*
261. 以图判树
给定编号从 0 到 n - 1 的 n 个结点。给定一个整数 n 和一个 edges 列表，其中 edges[i] = [ai, bi] 表示图中节点 ai 和 bi 之间存在一条无向边。
如果这些边能够形成一个合法有效的树结构，则返回 true ，否则返回 false 。
*/
func validTree(n int, edges [][]int) bool {
	graph := make([][]int, n)
	for i := range edges{
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
	}
	visited := make([]bool, n)
	var dfs func(u, p int)bool
	dfs = func(u, p int)bool{
		visited[u] = true
		for _, v := range graph[u]{
			if v == p{ continue }
			if !visited[v]{
				if !dfs(v, u){
					return false
				}
			}else{
				return false
			}
		}
		return true
	}
	//判断是否存在环，还有种情况， 多个分割的子图
	if !dfs(0, -1){
		return false
	}
	for i := range visited{
		if !visited[i]{
			return false
		}
	}
	return true
}
