package _200_2299


/*
2206. 将数组划分成相等数对
给你一个整数数组 nums ，它包含 2 * n 个整数。
你需要将 nums 划分成 n 个数对，满足：
每个元素 只属于一个 数对。
同一数对中的元素 相等 。
如果可以将 nums 划分成 n 个数对，请你返回 true ，否则返回 false 。
*/
func divideArray(nums []int) bool {
	var n [502]int
	for i:=0;i<len(nums);i++{
		n[nums[i]]++
	}
	for i:=0;i<len(n);i++{
		if n[i] %2 != 0 {
			return false
		}
	}
	return true
}
