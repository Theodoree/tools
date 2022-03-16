package _00_799



/*724. 寻找数组的中心索引*/
func pivotIndex(nums []int) int {
    var left, sum int
    for _, num := range nums {
        sum += num
    }
    for i, num := range nums {
        if left * 2 == sum - num {
            return i
        }
        left += num
    }
    return -1
}
