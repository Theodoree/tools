package __99

import "fmt"

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

func searchRange(nums []int, target int) []int {

    var left, right int
    result := []int{-1, -1}
    right = len(nums) - 1

    for left <= right {
        mid := (left + right) >> 1

        if nums[mid] > target {
            right = mid - 1
        } else if nums[mid] < target {
            left = mid + 1
        } else {

            for i := left; i <= right; i++ {
                if nums[i] == target {
                    result[0] = i
                    break
                }
            }

            for i := right; i >= left; i-- {
                if nums[i] == target {
                    result[1] = i
                    break
                }
            }
            break
        }
    }

    return result

}
func main() {

    f := searchRange([]int{2, 2}, 2)
    fmt.Println(f)

}
