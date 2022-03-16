package _00_199

//189. 旋转数组
// TODO 第一种解决方案
/*

func rotate(nums []int, k int) {
    len1 := len(nums)
    k = k % len1
    nums1 := make([]int, k)
    copy(nums1, nums[len1-k:])
    for i := len1 - k - 1; i >= 0; i-- {
        nums[i+k] = nums[i]
    }
    for j := 0; j < len(nums1); j++ {
        nums[j] = nums1[j]
    }
}

*/

//TODO 第二种解决方案 空间复杂度O(1)
func rotate(nums []int, k int) {


	for k > 0 {
		n:= nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = n

		k--
	}



}
// TODO 第三种解决方案

func main() {
    //rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
