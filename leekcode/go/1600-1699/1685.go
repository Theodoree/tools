package _600_1699




/*
1685. 有序数组中差绝对值之和
给你一个 非递减 有序整数数组 nums 。
请你建立并返回一个整数数组 result，它跟 nums 长度相同，且result[i] 等于 nums[i] 与数组中所有其他元素差的绝对值之和。
换句话说， result[i] 等于 sum(|nums[i]-nums[j]|) ，其中 0 <= j < nums.length 且 j != i （下标从 0 开始）。
*/
func getSumAbsoluteDifferences(nums []int) []int {
	preSum:=make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		if i == 0 {
			preSum[i] = nums[i]
			continue
		}
		preSum[i] = preSum[i-1] + nums[i]
	}

	var result []int
	for i:=0;i<len(nums);i++{

		if i == 0 {
			result = append(result,preSum[len(preSum)-1]-nums[i]*len(nums))
		}else{
			cnt:=nums[i]*(i+1) - preSum[i]
			if i != len(nums)-1 {
				cnt+=preSum[len(preSum)-1]-preSum[i] - nums[i]*(len(nums)-i-1)
			}
			result = append(result,cnt)
		}

	}

	return result

}

