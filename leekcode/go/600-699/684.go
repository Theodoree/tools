package _00_699

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

/*
684. 冗余连接
树可以看成是一个连通且 无环 的 无向 图。

给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。

请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。
*/
func findRedundantConnection(edges [][]int) []int {
	fs := newUfs(len(edges) + 1)

	for _, v := range edges {
		if !fs.isSame(v[0], v[1]) {
			fs.Union(v[0], v[1])
		} else {
			return v
		}

	}

	return nil
}
