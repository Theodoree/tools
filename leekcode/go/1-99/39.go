package __99

/*
39. 组合总和

给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/

func combinationSum(candidates []int, target int) [][]int {
    comb := [][]int{}
    for i := range candidates {
        if candidates[i] == target {
            comb = append(comb, []int{candidates[i]})
        } else if candidates[i] < target {
            rtn := combinationSum(candidates[i:], target-candidates[i])
            for j := range rtn {
                if len(rtn[j]) == 0 {
                    continue
                }
                comb = append(comb, append([]int{candidates[i]}, rtn[j]...))
            }
        }
    }
    return comb
}
