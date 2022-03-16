package __99

/*
45. 跳跃游戏 II
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。
*/
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var count int
	for i := 0; i < len(nums)-1; i++ {
		count++
		step := nums[i]
		idx := i + step
		max := i + step
		if max >= len(nums)-1 {
			break
		}
		for j := i + 1; j <= i+step && j < len(nums); j++ {
			if nums[j]+j > max {
				idx = j
				max = j + nums[j]
			}
		}
		if idx > i {
			i = idx - 1
		}
	}
	return count
}
