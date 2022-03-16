package _000_1099



/*
1060. 有序数组中的缺失元素
现有一个按 升序 排列的整数数组 nums ，其中每个数字都 互不相同 。
给你一个整数 k ，请你找出并返回从数组最左边开始的第 k 个缺失数字。
*/

func missingElement(nums []int, k int) int {
	start:=nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i] != start+1{
			sub:= nums[i]-start-1
			if k -sub <= 0 {
				return start+k
			}
			k-=sub
		}
		start = nums[i]
	}



	return start+k

}
