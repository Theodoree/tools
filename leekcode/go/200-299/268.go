package _00_299

import "fmt"

// 268. 缺失数字

func missingNumber(nums []int) int {

	array := make([]int, len(nums)+2)
	array[0] = -1
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if val <= len(nums) {
			array[val] = val
		}
	}

	for i := 0; i < len(array); i++ {

		if array[i] != i {
			return i
		}
	}

	return len(nums) + 1
}

func main() {

	f := missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	fmt.Println(f)

}
