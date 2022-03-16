package __99

//88. 合并两个有序数组

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j := 0, 0
	for {
		if i >= m+n || j >= n{
			break
		}
		if nums1[i] < nums2[j] && i < m + j {
			i++
		}else{
			nums1 = moveArray(nums1, i)
			nums1[i] = nums2[j]
			j++
		}
	}
}

func moveArray(nums1 []int, i int) []int{
	for j := len(nums1)-1; j > i; j-- {
		nums1[j] = nums1[j-1]
	}
	return nums1
}