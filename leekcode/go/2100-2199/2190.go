package _100_2199



/*
2190. 数组中紧跟 key 之后出现最频繁的数字
给你一个下标从 0 开始的整数数组 nums ，同时给你一个整数 key ，它在 nums 出现过。

统计 在 nums 数组中紧跟着 key 后面出现的不同整数 target 的出现次数。换言之，target 的出现次数为满足以下条件的 i 的数目：

0 <= i <= n - 2
nums[i] == key 且
nums[i + 1] == target 。
请你返回出现 最多 次数的 target 。测试数据保证出现次数最多的 target 是唯一的。
*/
func mostFrequent(nums []int, key int) int {
	var max int
	var maxNum int
	var buf [1002]int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			buf[nums[i+1]]++
			if buf[nums[i+1]] > max {
				max = buf[nums[i+1]]
				maxNum = nums[i+1]
			}
		}
	}

	return maxNum
}
