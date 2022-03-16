package _00_299


/*
229. 求众数 II

给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。

示例 1:

输入: [3,2,3]
输出: [3]
示例 2:

输入: [1,1,1,3,3,2,2,2]
输出: [1,2]
摩尔投票法
*/
func majorityElement(nums []int) []int {

    //最多二个众数
    var x, y, cx, cy int
    for i := 0; i < len(nums); i++ {
        if x == nums[i] {

            cx++
        } else if y == nums[i] {
            cy++
        } else if cx == 0 {  //防止重复
            x = nums[i]
            cx = 1
        } else if cy == 0 {
            y = nums[i]
            cy = 1
        } else {
            cy--
            cx--
        }
    }

    cx = 0
    cy = 0

    for i := 0; i < len(nums); i++ {
        if nums[i] == x {
            cx++
        } else if nums[i] == y {
            cy++
        }
    }
    var result []int
    if cx  > len(nums)/3{
        result = append(result,x)
    }
    if cy > len(nums)/3{
        result = append(result,y)
    }

    return  result
}