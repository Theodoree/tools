package _100_1199


/*
1137. 第 N 个泰波那契数

泰波那契序列 Tn 定义如下：

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。



示例 1：

输入：n = 4
输出：4
解释：
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4
示例 2：

输入：n = 25
输出：1389537
*/
func tribonacci(n int) int {

    switch n {
    case 0:
        return 0
    case 1:
        return 1
    case 2:
        return 1

    }

    s := make([]int, n+1)
    s[0] = 0
    s[1] = 1
    s[2] = 1
    for i := 3; i <= n; i++ {
        s[i] = s[i-1] + s[i-2] + s[i-3]
    }

    return s[n]
}
