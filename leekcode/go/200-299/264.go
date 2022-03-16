package _00_299




/*
264. 丑数 II

编写一个程序，找出第 n 个丑数。

丑数就是只包含质因数 2, 3, 5 的正整数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。
*/

func nthUglyNumber(n int) int {

    result := make([]int, 1)
    result[0] = 1

    var (
        l2 int
        l3 int
        l5 int
    )

    for i := 0; i < n-1; i++ {
        result = append(result, min(result[l2]*2, result[l3]*3, result[l5]*5))
        if result[len(result)-1] == result[l2]*2 {
            l2 += 1
        }
        if result[len(result)-1] == result[l3]*3 {
            l3 += 1
        }
        if result[len(result)-1] == result[l5]*5 {
            l5 += 1
        }
    }

    return result[len(result)-1]

}

func min(n ...int) int {

    var min int
    min = n[0]
    for _, v := range n {
        if v < min {
            min = v
        }
    }

    return min

}

