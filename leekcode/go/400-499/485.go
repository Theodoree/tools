package _00_499



/*
485. 最大连续1的个数

给定一个二进制数组， 计算其中最大连续1的个数。
*/
func findMaxConsecutiveOnes(nums []int) int {
    var max int
    var current int

    //for _,v:=range nums{
    for i := 0; i < len(nums); i++ {
        //if v == 1 {
        if nums[i] == 1 {
            current++
        } else {
            if current > max {
                max = current
            }

            current = 0
        }
    }

    if current > max {
        max = current
    }

    return max
}
