package _000_2099

/*
2006. 差的绝对值为 K 的数对数目
给你一个整数数组 nums 和一个整数 k ，请你返回数对 (i, j) 的数目，满足 i < j 且 |nums[i] - nums[j]| == k 。

|x| 的值定义为：

如果 x >= 0 ，那么值为 x 。
如果 x < 0 ，那么值为 -x 。

*/
func countKDifference(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {
			result := nums[j] - nums[i]
			if result == -k || result == k {
				count++
			}
		}

	}

	return count

}
