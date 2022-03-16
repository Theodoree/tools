package _900_1999

/*
1995. 统计特殊四元组
给你一个 下标从 0 开始 的整数数组 nums ，返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：
nums[a] + nums[b] + nums[c] == nums[d] ，且
a < b < c < d
*/
func countQuadruplets(nums []int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				result := nums[i] + nums[j] + nums[k]
				for val := k + 1; val < len(nums); val++ {
					if nums[val] == result {
						count++
					}
				}
			}
		}
	}
	return count
}
