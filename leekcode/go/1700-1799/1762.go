package _700_1799



/*
1762. 能看到海景的建筑物
有 n 座建筑物。给你一个大小为 n 的整数数组 heights 表示每一个建筑物的高度。
建筑物的右边是海洋。如果建筑物可以无障碍地看到海洋，则建筑物能看到海景。确切地说，如果一座建筑物右边的所有建筑都比它 矮 时，就认为它能看到海景。
返回能看到海景建筑物的下标列表（下标 从 0 开始 ），并按升序排列。
*/
func findBuildings(heights []int) []int {
	// 单调递增
	var result []int
	result = make([]int,len(heights))
	idx:=len(heights)-1
	result[idx] = len(heights)-1


	for i:=len(heights)-1;i>=0;i--{
		if heights[i] > heights[result[idx]] {
			idx--
			result[idx] = i
		}
	}

	return result[idx:]
}

