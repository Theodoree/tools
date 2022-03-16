package _00_199


/*
153. 寻找旋转排序数组中的最小值

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1:

输入: [3,4,5,1,2]
输出: 1
示例 2:

输入: [4,5,6,7,0,1,2]
输出: 0
*/
func findMin(nums []int) int {
    var min int
    min = math.MaxInt32
    if len(nums) < 5 {
        for _, v := range nums {
            if v < min {
                min = v
            }
        }
        return min
    }

    if nums[0] <nums[1] && nums[0] < nums[len(nums)-1]{
        return  nums[0]
    }

    right := len(nums) - 1

    for {
        if nums[right] > nums[right/2] {
            right /= 2
        } else {
            for i := right; i > 0; i-- {
                if nums[i-1] > nums[i] {
                    min = nums[i]
                    break
                }
            }
            break
        }
    }
    return min
}

/*


func findMin(nums []int) int {
    if len(nums) == 0{
        return -1
    } else if len(nums) == 1{
        return nums[0]
    }
    i:= 0
    j:= len(nums) - 1
    for i<j {
        mid := i + (j-i) / 2
        if (nums[mid] > nums[j]) {
            i = mid + 1
        } else {
            j = mid
        }

    }
    return nums[i]
}
*/