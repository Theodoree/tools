package _300_1399

/*
1319. 连通网络的操作次数
用以太网线缆将 n 台计算机连接成一个网络，计算机的编号从 0 到 n-1。线缆用 connections 表示，其中 connections[i] = [a, b] 连接了计算机 a 和 b。
网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。
给你这个计算机网络的初始布线 connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回 -1 。
*/

type ufs struct {
	roots []int
	cnt   int
}

func newUfs(n int) *ufs {
	roots := make([]int, n)
	for i := range roots {
		roots[i] = i
	}
	return &ufs{roots, n}
}
func (v *ufs) Find(x int) int {

	if v.roots[x] == x {
		return x
	} else {
		v.roots[x] = v.Find(v.roots[x])
		return v.roots[x]
	}

}

func (v *ufs) Union(x, y int) {
	v.roots[v.Find(x)] = v.Find(y)
}

func (v *ufs) isSame(x, y int) bool {
	return v.Find(x) == v.Find(y)
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	var result int
	u := newUfs(n)
	for _, v := range connections {
		u.Union(v[0], v[1])
	}

	for i := 0; i < n; i++ {
		if u.roots[i] == i {
			result++
		}
	}

	return result - 1
}
