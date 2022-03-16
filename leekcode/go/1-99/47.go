package __99

import (
	"fmt"
	"sort"
)

//47. 全排列 II

/*
47. 全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。



示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	PermuteUnique(nums, 0, len(nums)-1, &result)
	return result
}

func PermuteUnique(nums []int, left, right int, result *[][]int) {

	if left == right {
		*result = append(*result, nums)
		return
	}

	for i := left; i <= right; i++ {
		if i != left && nums[left] == nums[i] {
			continue
		}

		nums[left], nums[i] = nums[i], nums[left]
		buf := make([]int, len(nums))
		copy(buf[:], nums[:])
		PermuteUnique(buf, left+1, right, result)
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))

}
