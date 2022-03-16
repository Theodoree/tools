package _00_599



/*
523. 连续的子数组和
给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：
子数组大小 至少为 2 ，且
子数组元素总和为 k 的倍数。
如果存在，返回 true ；否则，返回 false 。
如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终视为 k 的一个倍数。
*/
func checkSubarraySum(nums []int, k int) bool {
	preSum:=make([]int,len(nums)+1)
	var zero bool
	var flag bool
	for i:=0;i<len(preSum);i++{
		if i == 0 {
			continue
		}
		if nums[i-1] == 0 {
			if zero{
				flag = true
			}
			zero = true
		}else{
			zero = false
		}
		preSum[i] = preSum[i-1] +nums[i-1]
	}
	if flag {
		return true
	}
	for i:=0;i<len(preSum);i++{
		if  preSum[len(preSum)-1] - preSum[i] < k {
			return false
		}

		for j:=i+2;j<len(preSum);j++{
			if (preSum[j] - preSum[i]) %k == 0 {
				return true
			}
		}
	}

	return false
}

