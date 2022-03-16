package _00_599

/*
533. 孤独像素 II
给你一个大小为 m x n 的二维字符数组 picture ，表示一张黑白图像，数组中的 'B' 表示黑色像素，'W' 表示白色像素。另给你一个整数 target ，请你找出并返回符合规则的 黑色 孤独像素的数量。
黑色孤独像素是指位于某一特定位置 (r, c) 的字符 'B' ，其中：
行 r 和列 c 中的黑色像素恰好有 target 个。
列 c 中所有黑色像素所在的行必须和行 r 完全相同。
*/
func findBlackPixel(picture [][]byte, N int) int {
	if 0 == len(picture) || 0 == len(picture[0]) {
		return 0
	}
	m, n := len(picture), len(picture[0])
	cs, ms := make([]int, n), map[string]int{}
	for i := 0; i < m; i++ {
		c := 0
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' {
				c++
				cs[j]++
			}
		}
		if c == N {
			ms[string(picture[i])]++
		}
	}
	out := 0
	for s, c := range ms {
		if c != N {
			continue
		}
		for i := 0; i < n; i++ {
			if s[i] == 'B' && cs[i] == N {
				out += N
			}
		}
	}
	return out
}

