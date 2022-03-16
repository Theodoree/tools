package __99


/*
81. 搜索旋转排序数组 II

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

示例 1:

输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
示例 2:

输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false
进阶:

这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
*/

func search(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }

    var (
        start int
        mid   = len(nums) / 2
        right = len(nums) - 1
    )

    for {
        if target == nums[start] || target == nums[mid] || target == nums[right] {
            return true
        }
        if start == mid {
            break
        }
        if nums[start] > nums[mid] {
            // 如果右边有序
            if target < nums[mid] || target > nums[right] {
                // 如果 target 小于中间，或者大于尾部，只可能在左侧
                right = mid
                mid = (start + right) / 2
            } else {
                // target 在右侧
                start = mid
                mid = (start + right) / 2
            }
        } else if nums[start] < nums[mid] {
            // 如果左边有序
            if target > nums[mid] || target < nums[start] {
                // 如果 target 大于中间，或者小于首部，只能在右侧
                start = mid
                mid = (start + right) / 2
            } else {
                // target 在左侧
                right = mid
                mid = (start + right) / 2
            }
        } else {
            // 如果头和中间相等，无法判断旋转点在哪边,去掉头和尾节点
            start++
            right--
        }
    }
    return false
}
