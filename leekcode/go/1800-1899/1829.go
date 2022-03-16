package _800_1899

/*
1829. 每个查询的最大异或值
给你一个 有序 数组 nums ，它由 n 个非负整数组成，同时给你一个整数 maximumBit 。你需要执行以下查询 n 次：
找到一个非负整数 k < 2maximumBit ，使得 nums[0] XOR nums[1] XOR ... XOR nums[nums.length-1] XOR k 的结果 最大化 。k 是第 i 个查询的答案。
从当前数组 nums 删除 最后 一个元素。
请你返回一个数组 answer ，其中 answer[i]是第 i 个查询的结果。
*/

func getMaximumXor(nums []int, maximumBit int) []int {
	preSum:=make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		if i == 0{
			preSum[i] = nums[i]
			continue
		}
		preSum[i] = preSum[i-1] ^ nums[i]
	}

	var max int
	for i:=0;i<maximumBit;i++{
		max |= 1<< i
	}

	var result []int
	for i:=0;i<len(nums);i++{
		s:=preSum[len(nums)-1-i]
		var n int
		for j:=0;j<maximumBit;j++{
			if s&(1<<j) == 0 {
				n+=1<<j
			}
		}
		result = append(result,n)

	}

	return result
}

