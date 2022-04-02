package sort

import (
	"golang.org/x/exp/constraints"
)

func SelectSort[T constraints.Integer](slice []T) []T {
	for index, _ := range slice {
		lowIndex := index
		for i := index + 1; i < len(slice); i++ {
			if slice[i] < slice[lowIndex] {
				lowIndex = i
			}
		}
		slice[index], slice[lowIndex] = slice[lowIndex], slice[index]
	}
	return slice
}

func InsertSort[T constraints.Integer](slice []T) []T {
	_InsertSort(slice, 0, len(slice)-1)
	return slice
}

func _InsertSort[T constraints.Integer](nums []T, left, right int) {
	for i := left + 1; i <= right; i++ {
		e := nums[i]
		for j := i; j > left && nums[j-1] > e; j-- {
			nums[j] = nums[j-1]
		}
	}

}

func BubbleSort[T constraints.Integer](slice []T) []T {
	for i := 0; i < len(slice); i++ {
		for j := i; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
