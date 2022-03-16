package __99

import "fmt"

// 35. 搜索插入位置

func searchInsert(nums []int, target int) int {


	switch  {
	case len(nums) == 0:
		return 0
	case target == 0:
		return 0
	case nums[0] > target:
		return  0
	}



	var index int

	for i := 0; i < len(nums); i++ {



		if nums[i] < target {
			index = i
		}

		if target == nums[i]{
			return  i
		}

	}

	return  index+1

}


func searchInsert1(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6},5))
	fmt.Println(searchInsert([]int{1,3,5,6},2))
	fmt.Println(searchInsert([]int{1,3,5,6},7))
	fmt.Println(searchInsert([]int{1,3,5,6},0))
	fmt.Println(searchInsert([]int{2,5},1))

}
