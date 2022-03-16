package _000_1099


/*
1099. 小于 K 的两数之和

给你一个整数数组 A 和一个整数 K，请在该数组中找出两个元素，使它们的和小于 K 但尽可能地接近 K，返回这两个元素的和。

如不存在这样的两个元素，请返回 -1。



示例 1：

输入：A = [34,23,1,24,75,33,54,8], K = 60
输出：58
解释：
34 和 24 相加得到 58，58 小于 60，满足题意。
示例 2：

输入：A = [10,20,30], K = 15
输出：-1
解释：
我们无法找到和小于 15 的两个元素。
*/

func twoSumLessThanK(A []int, K int) int {

    sort.Ints(A)

    var left, right int
    var max int
    max = -1
    right = len(A) - 1
    for left < right {

        sum := A[left] + A[right]

        if sum >= K {
            right--
        } else {

            max = int(math.Max(float64(max), float64(sum)))

            left++
        }

    }

    return max

}