package _500_1599

/*
1588. 所有奇数长度子数组的和
给你一个正整数数组 arr ，请你计算所有可能的奇数长度子数组的和。

子数组 定义为原数组中的一个连续子序列。

请你返回 arr 中 所有奇数长度子数组的和 。
*/
func sumOddLengthSubarrays(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var preSum = make([]int, len(arr)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + arr[i-1]
	}
	sum := preSum[len(preSum)-1]
	for i := 3; i < len(preSum); i += 2 {
		if i%2 == 1 {
			for j := i; j < len(preSum); j++ {
				sum += preSum[j] - preSum[j-i]
			}
		}
	}
	return sum

}
