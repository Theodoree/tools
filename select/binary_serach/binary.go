package binary_serach

//有序数组
func BinartySerach(array []int, target int) int {
	left := 0;
	right := len(array) - 1;
	for {
		mid := left + (right-left)/2
		if array[mid] == target {
			return mid
		}
		if target < array[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
}

//递归
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
