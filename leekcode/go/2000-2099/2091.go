package _000_2099



/*
2091. 从数组中移除最大值和最小值
给你一个下标从 0 开始的数组 nums ，数组由若干 互不相同 的整数组成。
nums 中有一个值最小的元素和一个值最大的元素。分别称为 最小值 和 最大值 。你的目标是从数组中移除这两个元素。
一次 删除 操作定义为从数组的 前面 移除一个元素或从数组的 后面 移除一个元素。
返回将数组中最小值和最大值 都 移除需要的最小删除次数。
*/
func minimumDeletions(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	minFn:= func(i,j int) int {
		if j < i {
			return j
		}
		return i
	}
	maxFn:= func(i,j int) int {
		if j > i {
			return j
		}
		return i
	}
	max:=math.MinInt64
	maxIdx:=-1
	min:=math.MaxInt64
	minIdx:=-1


	for i:=0;i<len(nums);i++{
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
		if nums[i] < min{
			min = nums[i]
			minIdx = i
		}
	}
	result:= minFn(maxFn(maxIdx+1,minIdx+1),maxFn(len(nums)-maxIdx,len(nums)-minIdx))
	result = minFn(result,maxIdx+1+len(nums)-minIdx)
	result = minFn(result,minIdx+1+len(nums)-maxIdx)

	return result
}

