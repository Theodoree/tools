package _go

import (
	"fmt"
	"sort"
	"strconv"
	"unsafe"
)

/*
剑指 Offer 45. 把数组排成最小的数
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
*/
// 30 < 3 < 34 < 5 < 9
// 9 > 5 > 34 > 3 > 30

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		return x < y
	})
	var result []byte
	for i := 0; i < len(nums); i++ {
		result = append(result, strconv.Itoa(nums[i])...)
	}

	return BytesToString(result)
}
