package _go


var dy = []int{0, -1, 0, 1}
var dx = []int{-1, 0, 1, 0}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// 染色 =》 转方向==》移动
func printKMoves(k int) (res []string) {
	if k == 0 {
		return []string{"R"}
	}
	var hash [4000][4000]bool
	si, sj, p := 2000, 2000, 3
	//hash[si][sj] = true
	minX, maxX, minY, maxY := si, sj, si, sj
	for tt := 0; tt < k; tt++ {
		//fmt.Println(si, sj, p)
		t := hash[si][sj]
		hash[si][sj] = !hash[si][sj] // 翻转
		if t {                       // 黑色，左逆时针90
			p = (p + 1) % 4
		} else {
			p = (p + 3) % 4
		}
		si, sj = si+dx[p], sj+dy[p]
		minX, minY = min(minX, si), min(minY, sj)
		maxX, maxY = max(maxX, si), max(maxY, sj)
	}
	//fmt.Println(maxX, minX, maxY, minY)
	bl := make([][]byte, maxX-minX+1)
	for i := minX; i <= maxX; i++ {
		bl[i-minX] = make([]byte, maxY-minY+1)
		for j := minY; j <= maxY; j++ {
			bl[i-minX][j-minY] = '_'
			if hash[i][j] {
				bl[i-minX][j-minY] = 'X'
			}
		}
	}
	bl[si-minX][sj-minY] = "ULDR"[p]
	for i := 0; i < len(bl); i++ {
		res = append(res, string(bl[i]))
	}
	return
}

