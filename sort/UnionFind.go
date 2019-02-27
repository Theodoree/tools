package sort

type UnionFind struct {
	Parent []int
	count  int
	Rank   []int
}

func NewUnionFind(n int) *UnionFind {
	var parent, rank []int

	for i := 0; i < n; i++ {
		parent = append(parent, i)
		rank = append(rank, i)
	}
	return &UnionFind{
		Parent: parent, count: n, Rank: rank,
	}
}

func (u *UnionFind) Size() int {
	return u.count
}

func (u *UnionFind) Find(p int) int {
	if p > 0 && u.count > p {
		for p != u.Parent[p] {
			p = u.Parent[p]
		}
		return p
	}
	return 0
}

//判断p和q的祖宗节点是否一致
func (u *UnionFind) IsConnected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UnionFind) UnionElements(p, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)

	if pRoot == qRoot {
		return
	}

	if u.Rank[pRoot] < u.Rank[qRoot] {
		u.Parent[pRoot] = qRoot
	} else if u.Rank[qRoot] < u.Rank[pRoot] {
		u.Parent[qRoot] = pRoot
	} else {
		u.Parent[pRoot] = qRoot
		u.Rank[qRoot]++
	}
}
