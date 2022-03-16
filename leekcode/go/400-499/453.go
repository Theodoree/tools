package _00_499


/*
453. 最小移动次数使数组元素相等

给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动可以使 n - 1 个元素增加 1。

示例:

输入:
[1,2,3]

输出:
3

解释:
只需要3次移动（注意每次移动会增加两个元素的值）：

[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
*/
func minMoves(nums []int) int {

    if len(nums) == 1 {
        return 0
    }

    var sum int
    var min int
    min = 1e9
    for _, v := range nums {
        if v < min {
            min = v
        }
        sum += v
    }
    return sum - len(nums)*min
}