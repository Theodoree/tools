package _400_1499

/*
1491. 去掉最低工资和最高工资后的工资平均值
给你一个整数数组 salary ，数组里每个数都是 唯一 的，其中 salary[i] 是第 i 个员工的工资。
请你返回去掉最低工资和最高工资以后，剩下员工工资的平均值。
*/
func average(salary []int) float64 {
	var sum, max, min int
	max = math.MinInt64
	min = math.MaxInt64
	for _, v := range salary {
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return float64(sum-max-min) / float64(len(salary)-2)
}
