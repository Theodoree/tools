package _00_599



/*
547. 省份数量
有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。
省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。
返回矩阵中 省份 的数量。
*/
func findCircleNum(isConnected [][]int) int {
	flag:=make([]bool,len(isConnected))
	cnt:=0

	idx:=make([]int,0,len(isConnected))

	for i:=0;i<len(isConnected);i++{
		if flag[i] {
			continue
		}
		cnt++
		idx = append(idx,i)
		bfs(isConnected,idx,flag)
		idx = idx[:0]
	}


	return cnt
}

func bfs(isConnected [][]int,idxs []int,flag []bool){
	for len(idxs) >  0 {
		idx:=idxs[len(idxs)-1]
		idxs = idxs[:len(idxs)-1]
		flag[idx]= true
		for i,v:=range isConnected[idx]{
			if flag[i] || i == idx{
				continue
			}
			if v == 1 {
				idxs = append(idxs, i)
			}
		}
	}
}


func findCircleNum(isConnected [][]int) int {
	// ufs
	u := newUfs(len(isConnected))
	for i := range isConnected {
		for j := i+1; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 {
				u.Union(i, j)
			}
		}
	}
	return u.Count()
}

type ufs struct {
	roots []int
	cnt int
}

func newUfs(n int) *ufs {
	roots := make([]int, n)
	for i := range roots {
		roots[i] = i
	}
	return &ufs{roots, n}
}
func(v *ufs) Find(x int) int {
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

func(v *ufs) Union(x, y int) {
	xRoot, yRoot := v.Find(x), v.Find(y)
	if xRoot == yRoot {
		return
	}
	v.cnt--
	v.roots[xRoot] = yRoot
}

func(v *ufs) Count() int {
	return v.cnt
}
