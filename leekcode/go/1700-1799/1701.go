package _700_1799


func averageWaitingTime(customers [][]int) float64 {
	start := customers[0][0]
	sum := 0
	for i := 0; i < len(customers); i++ {
		if customers[i][0] > start {
			start = customers[i][0]
		}
		start += customers[i][1]
		sum += start - customers[i][0]
	}
	return float64(sum) / float64(len(customers))
}

