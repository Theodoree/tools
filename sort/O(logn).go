package sort

import (
	"math"
	"math/rand"
)

func partition(slice []int,l,r int) int {

	_rand:=rand.Intn(r-l+1)+l
	slice[l],slice[_rand] = slice[_rand],slice[l]
	value:=slice[l]
	i:=l+1
	for{
		if i <=r && slice[i] < value{
			i++
		}

	}

}

func QuickSort(slice []int) []int {

	//if len(slice) < 16 {
	//	return InsertSort(slice)
	//}
	//
	//var leftSlice, rightSlice []int
	//d := rand.Intn(len(slice) - 1)
	//value := slice[d]
	//for _, v := range slice[1:] {
	//	if v <= value {
	//		leftSlice = append(leftSlice, v)
	//	} else {
	//		rightSlice = append(rightSlice, v)
	//	}
	//}
	//return append(append(QuickSort(leftSlice), value), QuickSort(rightSlice)...)
	quickSort(slice, 0, len(slice)-1)
	return slice
}

func quickSort(slice []int, start, end int) {

	if start > end {
		return
	}

	d := start+1
	value:=slice[0]
	j := end

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
	if (nums[mid] > nums[mid+1]) {
		_Merge(nums, left, mid, right)
	}
}

func _InsertSort(nums []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		e := nums[i]
		var j int
		for j := i; j > left && nums[j-1] > e; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = e
	}

}

func _Merge(nums []int, left, mid, right int) {
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

func MergeSortByDownToUp(slice []int) []int {
	for sz := 1; sz <= len(slice); sz += sz {
		for i := 0; i+sz < len(slice); i += sz * 2 {
			right := int(math.Min(float64(i+sz*2-1), float64(len(slice)-1)))
			if sz*2-sz < 16 {
				_InsertSort(slice, i, right)
			} else if slice[i+sz-1] > slice[i+sz] {
				_Merge(slice, i, i+sz-1, right)
			}
		}
	}
	return slice
}
