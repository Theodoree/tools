package _000_1099

import (
    "math"
    "sort"
)

/*
1046. 最后一块石头的重量

有一堆石头，每块石头的重量都是正整数。

每一回合，从中选出两块最重的石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。



提示：

1 <= stones.length <= 30
1 <= stones[i] <= 1000
*/
func lastStoneWeight(stones []int) int {

    if len(stones) == 1 {
        return stones[0]
    }
    if len(stones) == 0 {
        return -1
    }

    // n(logn)
    sort.Ints(stones)

    var maxIndex, secondIndex int
    maxIndex = len(stones) - 1
    secondIndex = len(stones) - 2
    for {
        v := stones[maxIndex] - stones[secondIndex]
        if v == 0 {
            stones[maxIndex] = 0
            stones[secondIndex] = 0
        } else {
            stones[maxIndex] = int(math.Abs(float64(v)))
            stones[secondIndex] = 0
        }
        // 这里使用INSERT算法  最坏的情况下两个元素都是零,则循环n*2次,最快的情况下循环n次
        sort.Ints(stones)

        if stones[secondIndex] == 0 {
            break
        }

    }
    return stones[maxIndex]
}

