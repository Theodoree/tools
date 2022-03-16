package _00_199

import (
    "fmt"
    "math"
)

// 121. 买卖股票的最佳时机
/*
记录【今天之前买入的最小值】
计算【今天之前最小值买入，今天卖出的获利】，也即【今天卖出的最大获利】
比较【每天的最大获利】，取最大值即可
*/
func maxProfit(prices []int) int {

    if len(prices) == 1 || len(prices) == 0 {
        return 0
    }

    min := float64(prices[0])
    var max float64

    for i := 1; i < len(prices); i++ {
        max = math.Max(float64(max), float64(prices[i])-min)
        min = math.Min(float64(min), float64(prices[i]))
    }

    return int(max)
}

func main() {

    fmt.Println(maxProfit([]int{1, 4, 2}))

}
