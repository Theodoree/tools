package _00_499

import (
    "math"
    "strconv"
)

/*
400. 第N个数字

在无限的整数序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...中找到第 n 个数字。

注意:
n 是正数且在32为整形范围内 ( n < 231)。

示例 1:

输入:
3

输出:
3
示例 2:

输入:
11

输出:
0

说明:
第11个数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是0，它是10的一部分。
*/

func findNthDigit(n int) int {
    var i int

    for n > i*int(math.Pow(10, float64(i-1))*9) {
        n -= i * int(math.Pow(10, float64(i-1))*9)
        i++
    }

    num := int(float64(n-1)/float64(i) + math.Pow(10, float64(i-1)))
    a := strconv.Itoa(num)
    if n%i == 0 {
        cnt, _ := strconv.Atoi(string(a[i-1]))
        return cnt
    }

    cnt, _ := strconv.Atoi(string(a[n%i-1]))
    return cnt
}
