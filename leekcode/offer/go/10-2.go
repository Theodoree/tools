package _go


/*
面试题10- II. 青蛙跳台阶问题

一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：

输入：n = 2
输出：2
示例 2：

输入：n = 7
输出：21
提示：

0 <= n <= 100
*/
func numWays(n int) int {
    switch n {
    case 0:
        return 1
    case 1:
        return 1
    case 2:
        return 2
    }
    var f = make([]int, n+1)

    f[0] = 1
    f[1] = 1

    for i := 2; i <= n; i++ {
        f[i] = f[i-1] + f[i-2]
        f[i] %= 1000000007
    }
    return f[n]
}
