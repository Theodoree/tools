package sort

import (
	"golang.org/x/exp/constraints"
	"math"
	"math/rand"
)

func QuickSort[T constraints.Integer](slice []T) []T {

	if len(slice) < 16 {
		return InsertSort(slice)
	}

	var leftSlice, rightSlice []T
	d := rand.Intn(len(slice) - 1)
	value := slice[d]
	slice[d], slice[0] = slice[0], slice[d]
	for _, v := range slice[1:] {
		if v <= value {
			leftSlice = append(leftSlice, v)
		} else {
			rightSlice = append(rightSlice, v)
		}
	}
	return append(append(QuickSort(leftSlice), value), QuickSort(rightSlice)...)
}

func QuickSort1[T constraints.Integer](slice []T) []T {
	var fn func(values []T, left, right int)
	fn = func(values []T, left, right int) {
		temp := values[left]
		p := left
		i, j := left, right

		for i <= j {
			for j >= p && values[j] >= temp {
				j--
			}
			if j >= p {
				values[p] = values[j]
				p = j
			}

			for i <= p && values[i] <= temp {
				i++
			}
			if i <= p {
				values[p] = values[i]
				p = i
			}
		}
		values[p] = temp
		if p-left > 1 {
			fn(values, left, p-1)
		}
		if right-p > 1 {
			fn(values, p+1, right)
		}
	}
	fn(slice, 0, len(slice)-1)
	return slice
}

func QuickSort2[T constraints.Integer](slice []T) []T {
	var fn func(slice []T)
	fn = func(slice []T) {
		if len(slice) <= 1 {
			return
		}
		mid, i := slice[0], 1
		head, tail := 0, len(slice)-1
		for head < tail {
			if slice[i] > mid {
				slice[i], slice[tail] = slice[tail], slice[i]
				tail--
			} else {
				slice[i], slice[head] = slice[head], slice[i]
				head++
				i++
			}
		}
		slice[head] = mid
		fn(slice[:head])
		fn(slice[head+1:])
	}
	fn(slice)
	return slice
}

func QuickSort3[T constraints.Integer](slice []T) []T {
	var fn func(slice []T, left int, right int)
	fn = func(slice []T, left int, right int) {
		if left >= right {
			return
		}

		explodeIndex := left

		for i := left + 1; i <= right; i++ {

			if slice[left] >= slice[i] {

				// 分割位定位++
				explodeIndex++
				slice[i], slice[explodeIndex] = slice[explodeIndex], slice[i]

			}

		}

		// 起始位和分割位
		slice[left], slice[explodeIndex] = slice[explodeIndex], slice[left]

		fn(slice, left, explodeIndex-1)
		fn(slice, explodeIndex+1, right)
	}

	fn(slice, 0, len(slice)-1)
	return slice
}

func MergeSort[T constraints.Integer](slice []T) []T {

	merge(slice, 0, len(slice)-1)
	return slice
}

func merge[T constraints.Integer](nums []T, left, right int) {

	if (right - left) < 16 {
		_InsertSort(nums, left, right)
		return
	}
	mid := (left + right) / 2
	merge(nums, left, mid)
	merge(nums, mid+1, right)
	if nums[mid] > nums[mid+1] {
		_merge(nums, left, mid, right)
	}
}

func _merge[T constraints.Integer](nums []T, left, mid, right int) {
	var NewNums []T
	for i := left; i <= right; i++ {
		NewNums = append(NewNums, nums[i])
	}
	var i = left    // 取起始位置
	var j = mid + 1 // 取中间位置

	for k := left; k <= right; k++ {
		if i > mid {
			nums[k] = NewNums[j-left]
			j++
		} else if j > right {
			nums[k] = NewNums[i-left]
			i++
		} else if NewNums[i-left] < NewNums[j-left] {
			nums[k] = NewNums[i-left]
			i++
		} else {
			nums[k] = NewNums[j-left]
			j++
		}
	}
}

func MergeSortByDownToUp[T constraints.Integer](slice []T) []T {
	for sz := 1; sz <= len(slice); sz += sz {
		for i := 0; i+sz < len(slice); i += sz * 2 {
			right := int(math.Min(float64(i+sz*2-1), float64(len(slice)-1)))
			if sz*2-sz < 16 {
				_InsertSort(slice, i, right)
			} else if slice[i+sz-1] > slice[i+sz] {
				_merge(slice, i, i+sz-1, right)
			}
		}
	}
	return slice
}
