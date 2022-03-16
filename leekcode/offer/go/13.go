package _go




func movingCount(m int, n int, k int) (res int) {
	visit := make([][]bool, m)
	for i, _ := range visit {
		visit[i] = make([]bool, n)
	}

	f1:= func(i int)  int {
		var cnt int
		for i> 0 {
			cnt+=i%10
			i/=10
		}
		return cnt
	}

	visit[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 || f1(i)+f1(j) > k {  // 起始点不用计算
				continue
			}
			if i >= 1 {  // 根据上一行是否能到达，来判断当前点是否可达
				visit[i][j] = visit[i][j] || visit[i-1][j]
			}
			if j >= 1 {   // 根据左一列是否能到达来判断
				visit[i][j] = visit[i][j] || visit[i][j-1]
			}
			if visit[i][j] {
				res++
			}
		}
	}
	return res + 1
}

