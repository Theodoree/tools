package _900_5999


/*
5996. 统计数组中相等且可以被整除的数对
给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 k ，请你返回满足 0 <= i < j < n ，nums[i] == nums[j] 且 (i * j) 能被 k 整除的数对 (i, j) 的 数目 。
*/
func countPairs(nums []int, k int) int {
	var cnt int
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i] == nums[j] && i*j%k == 0 {
				cnt++
			}
		}
	}
	return cnt
}

