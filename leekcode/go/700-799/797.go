package _00_799

import "runtime/debug"

/*
797. 所有可能的路径
给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）
二维数组的第 i 个数组中的单元都表示有向图中 i 号节点所能到达的下一些节点，空就是没有下一个结点了。
译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a 。
*/
func init() {
	debug.SetGCPercent(0)
}
func allPathsSourceTarget(graph [][]int) [][]int {

	var result [][]int
	result = make([][]int, 0, len(graph))
	var count int
	for _, v := range graph[0] {
		result = append(result, []int{0, v})
		if v == len(graph)-1 {
			count++
		}
	}

	for range graph {
		for i := 0; i < len(result); i++ {
			target := result[i][len(result[i])-1]
			if target == len(graph)-1 {
				continue
			}
			baseArr := result[i]
			for j := 0; j < len(graph[target]); j++ {
				if j == 0 {
					result[i] = append(result[i], graph[target][j])
				} else {
					buf := make([]int, 0, len(graph))
					buf = append(buf, baseArr...)
					buf = append(buf, graph[target][j])
					result = append(result, buf)
				}
				if graph[target][j] == len(graph)-1 {
					count++
				}
			}
		}
	}

	for i := 0; i < len(result); i++ {
		if result[i][len(result[i])-1] != len(graph)-1 {
			copy(result[i:], result[i+1:])
			result = result[:len(result)-1]
		}
	}

	return result
}
