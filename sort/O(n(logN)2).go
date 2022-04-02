package sort

import (
	"golang.org/x/exp/constraints"
)

func ShellSort[T constraints.Integer](slice []T) []T {
	//外层步长控制
	for step := len(slice) / 2; step > 0; step /= 2 {
		// 开始插入排序
		for i := step; i < len(slice); i++ {
			// 满足条件则插入
			for j := i - step; j >= 0 && slice[j+step] < slice[j]; j -= step {
				slice[j], slice[j+step] = slice[j+step], slice[j]
			}
		}
	}
	return slice
}
