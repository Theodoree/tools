package __99

// 27.移除元素

func removeElement(nums []int, val int) int {

	var count int
	for i := 0; i < len(nums); i++ {

		if nums[i] !=val {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}

