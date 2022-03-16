package _500_1599

/*
1566. 重复至少 K 次且长度为 M 的模式
给你一个正整数数组 arr，请你找出一个长度为 m 且在数组中至少重复 k 次的模式。
模式 是由一个或多个值组成的子数组（连续的子序列），连续 重复多次但 不重叠 。 模式由其长度和重复次数定义。
如果数组中存在至少重复 k 次且长度为 m 的模式，则返回 true ，否则返回  false 。
*/
func containsPattern(arr []int, m int, k int) bool {
	if m*k > len(arr) {
		return false
	}
	// 子数组,连续k次
	for i := 0; i <= len(arr)-m; i++ {
		subArr := arr[i : i+m]
		curK := k - 1
		curIdx := 0
		for j := i + m; j < len(arr); j++ {
			if arr[j] != subArr[curIdx] || curK == 0 {
				break
			}
			curIdx++
			if curIdx == len(subArr) {
				curIdx = 0
				curK--
			}
		}
		if curK <= 0 {
			return true
		}
	}
	return false

}
