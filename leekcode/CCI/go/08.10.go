package _go

/*
面试题 08.10. 颜色填充
编写函数，实现许多图片编辑软件都支持的「颜色填充」功能。
待填充的图像用二维数组 image 表示，元素为初始颜色值。初始坐标点的行坐标为 sr 列坐标为 sc。需要填充的新颜色为 newColor 。
「周围区域」是指颜色相同且在上、下、左、右四个方向上存在相连情况的若干元素。
请用新颜色填充初始坐标点的周围区域，并返回填充后的图像。
*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if newColor == image[sr][sc] {
		return image
	}
	bfs(image, sr, sc, newColor)
	return image
}
func bfs(image [][]int, sr, sc, newColor int) {

	ids := []int{sr, sc}
	old := image[sr][sc]
	image[sr][sc] = newColor
	for len(ids) > 0 {
		sr = ids[len(ids)-2]
		sc = ids[len(ids)-1]
		ids = ids[:len(ids)-2]

		// up
		if sr-1 >= 0 && image[sr-1][sc] == old {
			image[sr-1][sc] = newColor
			ids = append(ids, sr-1, sc)
		}

		// down
		if sr+1 < len(image) && image[sr+1][sc] == old {
			image[sr+1][sc] = newColor
			ids = append(ids, sr+1, sc)
		}

		// left
		if sc-1 >= 0 && image[sr][sc-1] == old {
			image[sr][sc-1] = newColor
			ids = append(ids, sr, sc-1)
		}

		// down
		if sc+1 < len(image[sr]) && image[sr][sc+1] == old {
			image[sr][sc+1] = newColor
			ids = append(ids, sr, sc+1)
		}

	}

}
