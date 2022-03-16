package _00_999


/*
930. 和相同的二元子数组
给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。
子数组 是数组的一段连续部分。
*/
func numSubarraysWithSum(nums []int, goal int) int {
	presum:=make([]int,len(nums)+1)
	for i:=1;i<len(presum);i++{
		presum[i] = presum[i-1] +nums[i-1]
	}
	var sum int
	idx:=0

	for i:=0;i<len(presum);i++{

		for presum[i] - presum[idx] > goal{
			idx++
		}

		var ok bool
		right := idx
		for right < i && right < len(presum) && presum[i] - presum[right] == goal{
			right++
			ok = true
		}

		if ok {
			sum+=right-idx
		}

	}

	return sum
}




func numSubarraysWithSum(nums []int, goal int) int {
	left1, left2 := 0, 0
	sum1, sum2 := 0, 0
	ans := 0
	// 滑窗，左边两个指针l1,l2 右边界为r
	// l1~r之间的和为s1，l2~r之间的和为s2
	// 对于每一个右边界r
	// 当l1~r之间的和大于goal时，令l1右移，直到区间和s1等于goal为止
	// （我想这是因为数组中的数字非0即1 所以区间和的步进只能是1，如果数组中还有别的数值，这里更新的过程可能需要修改）
	// 同理求取l2与l1之间的下标差值，这里停止l2右移的条件变成了区间和s2小于goal，即第一个s2不再大于等于goal的区间。
	// 意味着在以r为右边界，以l1为左边界的大区间内，以[l1~l2)之间任意下标为左边界，以r为右边界的小区间区间和都为goal
	// 符合区间之和等于goal的对应左下标有l2-l1个。
	// 统计符合条件的区间个数
	// r++ 更新新的右边界
	for right, v := range nums {
		sum1 += v
		for left1 <= right && sum1 > goal {
			sum1 -= nums[left1]
			left1++
		}
		sum2 += v
		for left2 <= right && sum2 >= goal {
			sum2 -= nums[left2]
			left2++
		}
		ans += left2 - left1
	}
	return ans
}