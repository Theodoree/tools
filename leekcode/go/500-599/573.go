package _00_599

/*
573. 松鼠模拟
现在有一棵树，一只松鼠和一些坚果。位置由二维网格的单元格表示。你的目标是找到松鼠收集所有坚果的最小路程，且坚果是一颗接一颗地被放在树下。松鼠一次最多只能携带一颗坚果，松鼠可以向上，向下，向左和向右四个方向移动到相邻的单元格。移动次数表示路程。
*/
func minDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	// 先计算所有nuts距离树的最近距离
	var sum []int
	for i := 0; i < len(nuts); i++ {
		row, col := nuts[i][0], nuts[i][1]
		sum = append(sum, (abs(tree[0]-row)+abs(tree[1]-col))*2)
	}
	var total int
	for _, v := range sum {
		total += v
	}

	min := math.MaxInt64
	for i := 0; i < len(sum); i++ {
		// 尝试从松鼠的位置坚果->树
		row, col := nuts[i][0], nuts[i][1]
		n := total - sum[i] + abs(squirrel[0]-row) + abs(squirrel[1]-col) + (abs(tree[0]-row) + abs(tree[1]-col))
		if n < min {
			min = n
		}
	}
	return min
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
