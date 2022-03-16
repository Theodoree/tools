package _00_299

import (
    "fmt"
)

func minCost(costs [][]int) int {

    if len(costs) == 0 {
        return 0
    }

    // 分析问题

    for i := 1; i < len(costs); i++ {
        costs[i][0] += Min(costs[i-1][1], costs[i-1][2])
        costs[i][1] += Min(costs[i-1][0], costs[i-1][2])
        costs[i][2] += Min(costs[i-1][0], costs[i-1][1])
    }

    fmt.Println(costs)

    var min int
    min = costs[len(costs)-1][0]
    for _, v := range costs[len(costs)-1] {
        min = Min(min, v)
    }

    return min

}

func Min(a, b int) int {
    if a > b {
        return b
    }

    return a

}