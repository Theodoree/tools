package _00_299



/*
263. 丑数

编写一个程序判断给定的数是否为丑数。

丑数就是只包含质因数 2, 3, 5 的正整数。

示例 1:

输入: 6
输出: true
解释: 6 = 2 × 3
示例 2:

输入: 8
输出: true
解释: 8 = 2 × 2 × 2
示例 3:

输入: 14
输出: false
解释: 14 不是丑数，因为它包含了另外一个质因数 7。
*/

func isUgly(num int) bool {

    if num <= 1 {
        if num == 1 {
            return true
        }
        return false
    }

    for num > 1 {
        t := num
        for t%2 == 0 {
            t /= 2
        }
        for t%3 == 0 {
            t /= 3
        }
        for t%5 == 0 {
            t /= 5
        }
        if t == num {
            return false
        }
        num = t
    }
    return true
}
