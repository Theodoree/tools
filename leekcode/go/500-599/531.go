package _00_599




/*
531. 孤独像素 I
给你一个大小为 m x n 的图像 picture ，图像由黑白像素组成，'B' 表示黑色像素，'W' 表示白色像素，请你统计并返回图像中 黑色 孤独像素的数量。
黑色孤独像素 的定义为：如果黑色像素 'B' 所在的同一行和同一列不存在其他黑色像素，那么这个黑色像素就是黑色孤独像素。
*/
func findLonelyPixel(picture [][]byte) int {
	if len(picture) == 0 {
		return 0
	}
	var col = make([]int, len(picture[0]))
	var row = make([]int, len(picture))
	for i := 0; i < len(picture); i++ {
		for j := 0; j < len(picture[i]); j++ {
			if picture[i][j] == 'B' {
				col[j] += 1 + i<<32
				row[i]++
			}
		}
	}

	var cnt int
	for i := 0; i < len(col); i++ {
		if col[i]&math.MaxInt32 == 1 && row[col[i]>>32] == 1 {
			cnt++
		}
	}

	return cnt
}
