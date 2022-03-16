package _00_199

/*
163. 缺失的区间
给定一个排序的整数数组 nums ，其中元素的范围在 闭区间 [lower, upper] 当中，返回不包含在数组中的缺失区间。
示例：
输入: nums = [0, 1, 3, 50, 75], lower = 0 和 upper = 99,
输出: ["2", "4->49", "51->74", "76->99"]
*/

func findMissingRanges(nums []int, lower int, upper int) []string {
	var result []string
	var left = math.MinInt64

	for i := lower; i <= upper; i++ {
		if len(nums) > 0 {

			n := nums[0]
			if n == i {
				nums = nums[1:]
				if left != math.MinInt64 {
					result = append(result, strconv.Itoa(left))
					left = math.MinInt64
				}
				continue
			}

			if left != math.MinInt64 {
				result = append(result, fmt.Sprintf("%d->%d", left, n-1))
				left = math.MinInt64
				i = n
				nums = nums[1:]
			} else {
				left = i
			}

		} else {
			if i == upper {
				result = append(result, strconv.Itoa(upper))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", i, upper))
			}
			break
		}

	}

	return result
}
