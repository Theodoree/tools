package sort




var target = 10

// 第一个值等于给定值的元素
func binary1(num []int) int {
    
    var (
        left  int
        right = len(num) - 1
    )
    
    for left <= right {
        
        mid := left + ((right - left) >> 1)
        
        if num[mid] > target {
            right = mid - 1
        } else if num[mid] < target {
            left = mid + 1
        } else {
            if left == 0 || num[mid-1] != target {
                return mid
            }
            left = mid - 1
        }
        
    }
    
    return -1
}

// 最后一个值等于给定值的元素
func binary2(num []int) int {
    var (
        left  int
        right = len(num) - 1
    )
    
    for left <= right {
        
        mid := left + ((right - left) >> 1)
        
        if num[mid] > target {
            right = mid - 1
        } else if num[mid] < target {
            left = mid + 1
        } else {
            if mid == len(num)-1 || num[mid+1] != target {
                return mid
            }
            left = mid + 1
        }
    }
    
    return -1
}

// 寻找第一个大于给定值的元素
func binary3(num []int) int {
    var (
        left  int
        right = len(num) - 1
    )
    
    for left <= right {
        
        mid := left + ((right - left) >> 1)
        
        if num[mid] >= target {
            
            if mid == 0 || num[mid-1] < target {
                for mid != len(num) && num[mid] == target {
                    mid++
                }
                return mid
            }
            right = mid - 1
        } else {
            left = mid + 1
        }
        
    }
    
    return -1
}

// 寻找第一个小于给定值的元素
func binary4(num []int) int {
    var (
        left  int
        right = len(num) - 1
    )
    
    for left <= right {
        
        mid := left + ((right - left) >> 1)
        
        if num[mid] > target {
            right = mid - 1
        } else {
            if mid == len(num)-1 || num[mid+1] > target {
                for mid > 0 && num[mid] == target {
                    mid--
                }
                if mid < 0 {
                    return -1
                }
                return mid
            }
            
            left = mid + 1
        }
        
    }
    
    return -1
}
// 有序数组
func BinartySerach(array []int, target int) int {
    left := 0
    right := len(array) - 1
    for {
        mid := left + (right-left)/2
        if array[mid] == target {
            return mid
        } else if target < array[mid] {
            right = mid - 1
        } else {
            left = mid + 1
        }
        
    }
}

// 递归
func BinartySerach_(array []int, left int, right int, target int) int {
    mid := left + (right-left)/2
    if target == array[mid] {
        return mid
    }
    if target < array[mid] {
        return BinartySerach_(array, left, mid-1, target)
    } else {
        return BinartySerach_(array, mid+1, right, target)
    }
}

func ceil(array []int, target int) int {
    left := -1
    right := len(array) - 1
    
    for left < right {
        mid := left + (right-left)/2
        if array[mid] > target {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    if right-1 >= 0 && array[right-1] == target {
        return right - 1;
    }
    return right;
}

func floor(array []int, target int) int {
    left := -1
    right := len(array) - 1
    
    for left < right {
        mid := left + (right-left)/2
        if array[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    
    if left+1 < len(array) && array[left+1] == target {
        return left + 1;
    }
    return left;
}
