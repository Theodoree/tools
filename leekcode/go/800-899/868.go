package _00_899



/*
868. 二进制间距

给定一个正整数 N，找到并返回 N 的二进制表示中两个连续的 1 之间的最长距离。

如果没有两个连续的 1，返回 0 。



示例 1：

输入：22
输出：2
解释：
22 的二进制是 0b10110 。
在 22 的二进制表示中，有三个 1，组成两对连续的 1 。
第一对连续的 1 中，两个 1 之间的距离为 2 。
第二对连续的 1 中，两个 1 之间的距离为 1 。
答案取两个距离之中最大的，也就是 2 。
示例 2：

输入：5
输出：2
解释：
5 的二进制是 0b101 。
示例 3：

输入：6
输出：1
解释：
6 的二进制是 0b110 。
示例 4：

输入：8
输出：0
解释：
8 的二进制是 0b1000 。
在 8 的二进制表示中没有连续的 1，所以返回 0 。


提示：

1 <= N <= 10^9
*/

func binaryGap(N int) int {

    var f = make([]int, 32)
    var cnt int
    for N > 0 {
        f[cnt] = N & 1
        N >>= 1
        cnt++
    }

    var max int
    var first, secord int
    for first < len(f) {
        for first < len(f)-1 && f[first] != 1 {
            first++
        }
        secord = first + 1
        if secord >= len(f) {
            break
        }

        for secord < len(f)-1 && f[secord] != 1 {
            secord++
        }

        if secord >= len(f) {
            break
        }

        if f[secord] == 1 && f[first] == 1 {

            if secord-first > max {
                max = secord - first
            }
        }

        first = secord

    }
    return max
}
