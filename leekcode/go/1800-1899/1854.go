package _800_1899

/*
1854. 人口最多的年份
给你一个二维整数数组 logs ，其中每个 logs[i] = [birthi, deathi] 表示第 i 个人的出生和死亡年份。
年份 x 的 人口 定义为这一年期间活着的人的数目。第 i 个人被计入年份 x 的人口需要满足：x 在闭区间 [birthi, deathi - 1] 内。注意，人不应当计入他们死亡当年的人口中。
返回 人口最多 且 最早 的年份。
*/
func maximumPopulation(logs [][]int) int {
	var buf [101]int
	var death [101]int
	for _, v := range logs {
		buf[v[0]-1950]++
		death[v[1]-1950]++
	}

	var maxCount int
	var maxI int

	var count int
	for i := 1; i < 100; i++ {
		count += buf[i] - death[i]
		if count > maxCount {
			maxCount = count
			maxI = i
		}
	}

	return maxI + 1950

}
