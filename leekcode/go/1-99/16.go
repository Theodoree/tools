package __99

import (
    "math"
    "sort"
)
/*
16. 最接近的三数之和

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/

func threeSumClosest(nums []int, target int) int {
    if len(nums) == 0  {
        return 0
    }
    sort.Ints(nums)

    var left,right int
    closenum := nums[0] + nums[1] + nums[2]
    for i := 0; i < len(nums)-2; i++ {
        if nums[i]+nums[i+1] > target {
            break
        }
        left = i + 1
        right = len(nums) - 1
        for left < right {

            threeSum := nums[i] + nums[left] + nums[right]
            if math.Abs(float64(threeSum-target)) < math.Abs(float64(closenum-target)) {
                closenum = threeSum
            }

            if threeSum > target {
                right--
            } else if threeSum < target {
                left++
            } else {
                return threeSum
            }
        }

    }
    return closenum
}