package _00_599

import (
    "fmt"
    "sort"
)

/*
506. 相对名次

给出 N 名运动员的成绩，找出他们的相对名次并授予前三名对应的奖牌。前三名运动员将会被分别授予 “金牌”，“银牌” 和“ 铜牌”（"Gold Medal", "Silver Medal", "Bronze Medal"）。

(注：分数越高的选手，排名越靠前。)

示例 1:

输入: [5, 4, 3, 2, 1]
输出: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
解释: 前三名运动员的成绩为前三高的，因此将会分别被授予 “金牌”，“银牌”和“铜牌” ("Gold Medal", "Silver Medal" and "Bronze Medal").
余下的两名运动员，我们只需要通过他们的成绩计算将其相对名次即可。
提示:

N 是一个正整数并且不会超过 10000。
所有运动员的成绩都不相同。

*/

func findRelativeRanks(nums []int) []string {
    First := "Gold Medal"
    Second := "Silver Medal"
    Third := "Bronze Medal"

    rank:=make([]int,len(nums))
    for k,_:=range nums{
        rank[k] = k
    }
    sort.Slice(rank, func(i, j int) bool {
        return nums[rank[i]] > nums[rank[j]]
    })


    result := make([]string, len(nums))
    for k := 0; k < len(nums); k++ {
        switch k {
        case 0:
            result[rank[0]] = First
        case 1:
            result[rank[1]] = Second
        case 2:
            result[rank[2]] = Third
        default:
            result[rank[k]] = fmt.Sprintf("%d", k+1)
        }

    }

    return result
}

var medals = []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(nums []int) []string {
    index := map[int]int{}
    for i, n := range nums {
        index[n] = i
    }
    sort.Ints(nums)
    res := make([]string, len(nums))

    for i := 0; i < len(nums); i++ {
        j := len(nums) - 1 - i
        k := index[nums[j]]
        if i < 3 {
            res[k] = medals[i]
        } else {
            res[k] = fmt.Sprintf("%d", i+1)
        }
    }

    return res
}

