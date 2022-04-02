package sort

import (
	"golang.org/x/exp/constraints"
)

// 第一个值等于给定值的元素

func BinarySearch[T constraints.Integer](num []T, target T) int {

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

func BinarySearchLast[T constraints.Integer](num []T, target T) int {
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

func BinarySearchBigger[T constraints.Integer](num []T, target T) int {
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

func BinarySearchLess[T constraints.Integer](num []T, target T) int {
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
