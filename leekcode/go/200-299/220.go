package _00_299


func Quick3Sort(slice []int,idx []int) []int {
	quick3Sort(slice, idx,0, len(slice)-1)
	return slice
}
func quick3Sort(a []int, idx []int,left int, right int) {

	if left >= right {
		return
	}

	explodeIndex := left

	for i := left + 1; i <= right; i++ {

		if a[left] >= a[i] {

			// 分割位定位++
			explodeIndex++;
			idx[i],idx[explodeIndex] = idx[explodeIndex],idx[i]
			a[i], a[explodeIndex] = a[explodeIndex], a[i]

		}

	}

	// 起始位和分割位
	idx[left],idx[explodeIndex] = idx[explodeIndex],idx[left]
	a[left], a[explodeIndex] = a[explodeIndex], a[left]

	quick3Sort(a, idx,left, explodeIndex-1)
	quick3Sort(a, idx,explodeIndex+1, right)
}

/*
220. 存在重复元素 III
给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
如果存在则返回 true，不存在返回 false。
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	abs:= func(i int) int {
		if i< 0 {
			return -i
		}
		return i
	}
	var idx =make([]int,len(nums))
	for i:=0;i<len(idx);i++{
		idx[i] = i
	}

	Quick3Sort(nums,idx)


	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[j] - nums[i] > t {
				break
			}
			if abs(idx[j] - idx[i]) <= k{
				return true
			}
		}
	}


	return false
}

