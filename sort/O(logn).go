package sort

import (
	"math"
	"math/rand"
)

func QuickSort(slice []int) []int {

	if len(slice) < 16 {
		return InsertSort(slice)
	}

	var leftSlice, rightSlice []int
	d := rand.Intn(len(slice) - 1)
	value := slice[d]
	for _, v := range slice[1:] {
		if v <= value {
			leftSlice = append(leftSlice, v)
		} else {
			rightSlice = append(rightSlice, v)
		}
	}
	return append(append(QuickSort(leftSlice), value), QuickSort(rightSlice)...)
}

func MergeSort(slice []int) []int {

	merge(slice, 0, len(slice)-1)
	return slice
}

func merge(nums []int, left, right int) {

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

func MergeSortByDownToUp(slice []int) []int {
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

func _merge(nums []int, left, mid, right int) {
	var NewNums []int
	for i := left; i <= right; i++ {
		NewNums = append(NewNums, nums[i])
	}
	var i = left    //取起始位置
	var j = mid + 1 //取中间位置

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

func _InsertSort(nums []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		e := nums[i]
		for j := i; j > left && nums[j-1] > e; j-- {
			nums[j] = nums[j-1]
		}
	}

}

func HeapSort(array []int) []int {

	array = _heapInsert(array)
	//_heapShiftUp(array)
	//_heapShiftDown(array)

	for i := 0; i < len(array); i++ {
		array[i] = _heapExtractMax(array)
	}

	return array
}

func _heapInsert(array []int) []int {

	var newArray []int

	for i := 0; i < len(array); i++ {
		newArray = append(newArray, array[i])
		lastIndex := len(newArray) - 1
		for newArray[lastIndex/2] > newArray[lastIndex] {
			newArray[lastIndex/2], newArray[lastIndex] = newArray[lastIndex], newArray[lastIndex/2]
			lastIndex /= 2
		}
	}
	return newArray
}

func _heapExtractMax(array []int) int {
	array[1], array[len(array)-1] = array[len(array)-1], array[1]
	_heapShiftDown(array)
	return 0
}
func _heapShiftUp(array []int, value int) {
	array = append(array, value)
	index := len(array) - 1
	for index >= 1 && array[index/2] < array[index] {
		array[index/2], array[index] = array[index], array[index/2]
		index /= 2;
	}
}

func _heapShiftDown(array []int) (slice []int, value int) {
	value = array[0]
	array[0] = array[len(array)-1]
	array = array[:len(array)-1]
	var index int
	for len(array) > index*2+1 || len(array) > index*2 {
		var left, right int
		if len(array) > index*2 && array[index*2] > left {
			left = array[index*2]
		}

		if len(array) > index*2+1 && array[index*2+1] > right {
			right = array[index*2+1]
		}

		if left > right {
			array[index], array[index*2] = array[index*2], array[index]
			index *= 2
		} else {
			array[index], array[index*2+1] = array[index*2+1], array[index]
			index = index*2 + 1
		}
	}
	return array, value
}
