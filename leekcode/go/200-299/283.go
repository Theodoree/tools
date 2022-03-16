package _00_299



//移动零
func moveZeroes(nums []int) {
    o := 0
    for i := 0; i < len(nums); i++ {
        v := nums[i]
        if v == 0 {
            continue
        } else {
            nums[o] = v
            o++
        }
    }
    for i := o; i < len(nums); i++ {
        nums[i] = 0
    }
}
