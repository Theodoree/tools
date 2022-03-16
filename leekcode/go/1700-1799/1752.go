package _700_1799

/*
1752. 检查数组是否经排序和轮转得到
给你一个数组 nums 。nums 的源数组中，所有元素与 nums 相同，但按非递减顺序排列。
如果 nums 能够由源数组轮转若干位置（包括 0 个位置）得到，则返回 true ；否则，返回 false 。
源数组中可能存在 重复项 。
注意：我们称数组 A 在轮转 x 个位置后得到长度相同的数组 B ，当它们满足 A[i] == B[(i+x) % A.length] ，其中 % 为取余运算。
*/
func check(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			cnt := i
			for cnt < cnt+len(nums) {
				if nums[cnt%len(nums)] > nums[(cnt+1)%len(nums)] {
					break
				}
				cnt++
			}
			return cnt == i+len(nums)-1
		}
	}

	return true
}
