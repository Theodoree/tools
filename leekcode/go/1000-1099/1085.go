package _000_1099

import "math"

/*
1085. 最小元素各数位之和

给你一个正整数的数组 A。

然后计算 S，使其等于数组 A 当中最小的那个元素各个数位上数字之和。

最后，假如 S 所得计算结果是 奇数 的请你返回 0，否则请返回 1。



示例 1:

输入：[34,23,1,24,75,33,54,8]
输出：0
解释：
最小元素为 1，该元素各个数位上的数字之和 S = 1，是奇数所以答案为 0。
示例 2：

输入：[99,77,33,66,55]
输出：1
解释：
最小元素为 33，该元素各个数位上的数字之和 S = 3 + 3 = 6，是偶数所以答案为 1。
*/

func sumOfDigits(A []int) int {

    var min int
    min = math.MaxInt32
    for _, v := range A {
        if v < min {
            min = v
        }
    }

    var cnt int

    for min > 0 {
        cnt += min % 10
        min /= 10
    }
    if cnt%2 == 0 {
        return 1
    }

    return 0

}