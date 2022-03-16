package _00_799

import "math"

/*
746. 使用最小花费爬楼梯
数组的每个索引做为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。

每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。

您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

*/

func minCostClimbingStairs(cost []int) int {

    switch {
    case len(cost) == 2:
        if cost[0] > cost[1] {
            return cost[1]
        } else {
            return cost[0]
        }
    case len(cost) == 1:
        return cost[0]
    case len(cost) == 0:
        return 0
    }

    a := cost[0]
    b := cost[1]
    var c int
    for i := 2; i < len(cost); i++ {
        if a > b {
            c = b
        } else {
            c = a
        }

        c += cost[i]
        a = b
        b = c
    }
    return  int(math.Min(float64(a),float64(b)))
}