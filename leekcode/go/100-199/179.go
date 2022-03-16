package _00_199

import (
	"sort"
	"strconv"
	"strings"
)

/*
179. 最大数
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"
示例 3：

输入：nums = [1]
输出："1"
示例 4：

输入：nums = [10]
输出："10"


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 109

*/

func largestNumber(nums []int) string {

	/*
	实质的逻辑在于该题是要求拼接最大数,那么数字的排序,只需要知道
	xy  > yx OR yx > xy

	*/
	sort.Slice(nums, func(i, j int) bool {
		i1:=10
		j1:=10

		for i1 <=nums[i] {
			i1 *=10
		}
		for j1 <=nums[j] {
			j1 *=10
		}

		return j1 * nums[i] + nums[j]  >  i1 *  nums[j] + nums[i]
	})



	var sum strings.Builder
	var zero bool
	zero = true
	for _,v:=range nums{
		if v != 0 && zero{
			zero = false
		}
		sum.WriteString(strconv.Itoa(v))
	}


	if zero {
		return "0"
	}

	return sum.String()
}

