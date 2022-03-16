package _00_299

import "math"

/*
233. 数字 1 的个数

给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。

示例:

输入: 13
输出: 6
解释: 数字 1 出现在以下数字中: 1, 10, 11, 12, 13 。

*/

func countDigitOne(n int) int {
    num := n
    i := 1
    s := 0

    for num > 0 {
        if num%10 == 0 {
            s += (num / 10) * i
        }
        if num%10 == 1 {
            s += (num/10)*i + (n % i) + 1
        }
        if num%10 > 1 {
            s += int(math.Ceil(float64(num) / 10.0)) * i
        }
        num = num / 10
        i = i * 10

    }
    return s
}

