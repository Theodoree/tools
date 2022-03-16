package _000_2099

/*
2057. 值相等的最小索引
给你一个下标从 0 开始的整数数组 nums ，返回 nums 中满足 i mod 10 == nums[i] 的最小下标 i ；如果不存在这样的下标，返回 -1 。

x mod y 表示 x 除以 y 的 余数 。
*/

func smallestEqual(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i%10 {
			return i
		}
	}
	return -1
}
