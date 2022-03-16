package __99

import (
    "sort"
)

/*
40. 组合总和 II

给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
*/

func combinationSum2(candidates []int, target int) [][]int {

    sort.Ints(candidates)
    var result [][]int
    CombinationSum2(candidates, &result, []int{}, 0, target)
    return result
}

func CombinationSum2(nums []int, result *[][]int, cur []int, level int, target int) {

    if target == 0 {
        var c = make([]int, len(cur))
        copy(c, cur)
        *result = append(*result, c)
        return
    } else if target < 0 {
        return
    }

    for i := level; i < len(nums); i++ {

        if target-nums[i] >= 0 { // 如果小于0肯定不行
            if i > level && nums[i] == nums[i-1] {
                continue
            }
            cur = append(cur, nums[i])
            CombinationSum2(nums, result, cur, i+1, target-nums[i])
            cur = cur[0 : len(cur)-1]
        } else {
            return
        }
    }

}
