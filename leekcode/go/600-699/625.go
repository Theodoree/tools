package _00_699

import "math"

/*
625. 最小因式分解

给定一个正整数 a，找出最小的正整数 b 使得 b 的所有数位相乘恰好等于 a。

如果不存在这样的结果或者结果不是 32 位有符号整数，返回 0。

样例 1

输入：48
输出：68


样例 2

输入：15
输出：35



10
*/

func smallestFactorization(a int) int {
    if a < 10 {
        return a
    }
    var cnt, result int
    cnt = 1
    for i := 9; i >= 2; i-- {

        for a%i == 0 {
            result += cnt * i
            if result > math.MaxInt32 {
                return 0
            }
            a /= i
            cnt *= 10
        }
    }
    if a == 1 {
        return result
    }

    return 0

}
