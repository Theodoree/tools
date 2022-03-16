package _200_1299

/*
1273. 删除树节点
给你一棵以节点 0 为根节点的树，定义如下：
节点的总数为 nodes 个；
第 i 个节点的值为 value[i] ；
第 i 个节点的父节点是 parent[i] 。
请你删除节点值之和为 0 的每一棵子树。
在完成所有删除之后，返回树中剩余节点的数目。
*/
func deleteTreeNodes(nodes int, parent []int, value []int) int {
	var m = make(map[int][]int, nodes)
	for i := 0; i < len(value); i++ {
		m[parent[i]] = append(m[parent[i]], i)
	}
	dfs(0, m, value, &nodes)

	return nodes
}

func dfs(start int, m map[int][]int, value []int, nodes *int) (int, int, bool) {
	sum := value[start]
	child := m[start]
	count := len(child)
	for _, v := range child {
		vals, cnt, del := dfs(v, m, value, nodes)
		sum += vals
		count += cnt
		if del {
			count--
		}
	}
	if sum == 0 {
		*nodes -= count + 1
		delete(m, start)
		return 0, 0, true
	}
	return sum, count, false
}
