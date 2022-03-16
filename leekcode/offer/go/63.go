package _go

import "math"

/*
剑指 Offer 63. 股票的最大利润
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
*/
func maxProfit(prices []int) int {

	if len(prices) == 1 || len(prices) == 0 {
		return 0
	}

	min := float64(prices[0])
	var max float64

	for i := 1; i < len(prices); i++ {
		max = math.Max(float64(max), float64(prices[i])-min)
		min = math.Min(float64(min), float64(prices[i]))
	}

	return int(max)
}
