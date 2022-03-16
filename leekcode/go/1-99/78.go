package __99

/*
78. 子集

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

func subsets(nums []int) [][]int {
    var result [][]int
    var l uint = uint(len(nums))
    var n = 1 << l
    for i := 0; i < n; i++ {
        var cur []int
        for j := 0; j < int(l); j++ {
            if i&(1<<uint(j)) == 0 { // 这里强转是因为leekcode的go版本太低,高版本已支持负位运算,所以可以使用int类
                cur = append(cur, nums[j])
            }
        }
        result = append(result, cur)
    }

    return result
}

func main() {
    subsets([]int{1, 2, 3})
}
