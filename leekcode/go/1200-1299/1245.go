package _200_1299




/*
1245. 树的直径
给你这棵「无向树」，请你测算并返回它的「直径」：这棵树上最长简单路径的 边数。
我们用一个由所有「边」组成的数组 edges 来表示一棵无向树，其中 edges[i] = [u, v] 表示节点 u 和 v 之间的双向边。
树上的节点都已经用 {0, 1, ..., edges.length} 中的数做了标记，每个节点上的标记都是独一无二的。
*/
func treeDiameter(edges [][]int) int {
	var m = make(map[int][]int)
	for _, v := range edges {
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}

	var max int
	var idx int
	dfs(m,0,0,0,&max,&idx)

	dfs(m,idx,idx,0,&max,&idx)

	return max
}
func dfs(edges map[int][]int, parent int, cur int,depth int, max *int,idx *int)  {
	if depth > *max {
		*max = depth
		*idx = cur
	}
	for _, v := range edges[cur] {
		if v == parent {
			continue
		}
		dfs(edges, cur, v,depth+1, max,idx)
	}

	return
}

