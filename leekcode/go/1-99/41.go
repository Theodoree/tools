package __99

import "fmt"

//41. 缺失的第一个正数

func firstMissingPositive(nums []int) int {

	if len(nums) == 0 {
		return 1
	}

	array := make([]int, len(nums)+5)
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if val > 0 && val <= len(nums) {
			array[val] = val
		}
	}

	for i := 1; i < len(array); i++ {

		if array[i] != i {
			return i
		}

	}

	return len(array) + 1
}

func main() {
	f := firstMissingPositive([]int{1})
	fmt.Println(f)

}
