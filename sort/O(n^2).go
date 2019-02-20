package sort

func SelectSort(arr []int) []int {
	for index, _ := range arr {

		lowIndex := index

		for i := index + 1; i < len(arr); i++ {
			if arr[i] < arr[lowIndex] {
				lowIndex = i
			}
		}
		arr[index], arr[lowIndex] = arr[lowIndex], arr[index]
	}
	return arr
}

func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			j := i - 1
			temp := nums[i]
			for j >= 0 && nums[j] > temp {
				nums[j+1] = nums[j]
				j--
			}
			nums[j+1] = temp
		}
	}
	return nums
}

