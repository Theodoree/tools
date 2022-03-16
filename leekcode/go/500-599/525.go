package _00_599


/*
525. 连续数组
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
*/

func findMaxLength(nums []int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            nums[i] = -1
        }
    }

    m := make(map[int]int)
    m[0] = -1
    var result int
    var sum int
    for i := 0; i < len(nums); i++ {
        sum+=nums[i]
        if val,ok:=m[sum];ok{
            if result < i-val{
                result = i-val
            }
        }else{
            m[sum] = i
        }

    }

    return result
}