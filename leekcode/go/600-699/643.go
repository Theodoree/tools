package _00_699


/*
643. 子数组最大平均数 I

给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。

示例 1:

输入: [1,12,-5,-6,50,3], k = 4
输出: 12.75
解释: 最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
*/
func findMaxAverage(nums []int, k int) float64 {
    var max int
    var sum int
    for _, v := range nums[:k] {
        max += v
    }

    sum=max
    for i := 1; i <= len(nums)-k; i++ {

        sum = sum - nums[i-1] + nums[k+i-1]
        if sum > max {
            max = sum
        }


    }

    return float64(max) / float64(k)


}
