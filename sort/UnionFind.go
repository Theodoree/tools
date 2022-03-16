package sort

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
	root := x
	for root != v.roots[root] {
		root = v.roots[root]
	}
	for root != v.roots[x] {
		cur := v.roots[x]
		v.roots[x] = root
		x = cur
	}
	return root
}

func (v *ufs) Union(x, y int) {
	xRoot, yRoot := v.Find(x), v.Find(y)
	if xRoot == yRoot {
		return
	}
	v.cnt--
	v.roots[xRoot] = yRoot
}

func (v *ufs) Count() int {
	return v.cnt
}
