package _00_299

import "fmt"

//287. 寻找重复数 不能排序 只能使用额外O(1)的空间
//TODO 双指针
func findDuplicate(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	first := 0
	last := len(nums) - 1

	for {

		if first == last {
			first++
			last = len(nums) - 1
		}

		if nums[first] == nums[last] {
			return nums[first]
		}

		last--
	}

}

/*  TODO 遍历
func findDuplicate(nums []int) int {
	for _, v := range nums {
		i := abs(v)
		if nums[i] < 0 {
			return i
		} else {
			nums[i] = -nums[i]
		}
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
*/
func main() {

	f := findDuplicate([]int{1, 3, 4, 2, 2})
	fmt.Println(f)

}
