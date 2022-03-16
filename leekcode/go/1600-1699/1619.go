package _600_1699

import "sort"

/*
1619. 删除某些元素后的数组均值
给你一个整数数组 arr ，请你删除最小 5% 的数字和最大 5% 的数字后，剩余数字的平均值。
与 标准答案 误差在 10-5 的结果都被视为正确结果。
*/
func trimMean(arr []int) float64 {
	sort.Ints(arr)
	var sum int
	n := int(float64(len(arr)) * 0.05)

	for i := n; i < len(arr)-n; i++ {
		sum += arr[i]
	}

	return float64(sum) / float64(len(arr)-n*2)

}
