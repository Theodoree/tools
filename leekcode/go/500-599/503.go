package _00_599


/*
503. 下一个更大元素 II

给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

示例 1:

输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
注意: 输入数组的长度不会超过 10000。
*/

type stack struct {
    next *stack
    val  int
    idx  int
}

func nextGreaterElements(nums []int) []int {
    var s *stack
    
    var arr = make([]int, len(nums))
    var current int
    
    for i := len(nums) - 1; i >= 0; i-- {
        f := &stack{next: s, val: nums[i], idx: i}
        s = f
    }
    
    cur := s
    for current < len(nums) {
        c := cur
        for c != nil && nums[current] >= c.val {
            c = c.next
        }
        if c == nil {
            c := s
            for c.idx < current && c.val <= nums[current] {
                c = c.next
            }
            if c.val > nums[current] {
                arr[current] = c.val
            } else {
                arr[current] = -1
            }
        } else {
            arr[current] = c.val
        }
        cur = cur.next
        current++
    }
    
    return arr
    
}
