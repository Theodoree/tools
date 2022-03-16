package _00_699

import (
    "math"
)


/*633. 平方数之和
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c。

示例1:

输入: 5
输出: True
解释: 1 * 1 + 2 * 2 = 5


示例2:

输入: 3
输出: False
*/

func judgeSquareSum(c int) bool {
    j := int(math.Sqrt(float64(c)))
    i := 0

    for i <= j {
        total := i*i + j*j
        if total > c {
            j -= 1
        } else if total < c {
            i += 1
        } else {
            return true
        }
    }
    return false

}
