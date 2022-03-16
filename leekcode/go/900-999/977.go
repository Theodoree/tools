package _00_999

import (
	"sort"
)

// 977. 有序数组的平方

func sortedSquares(A []int) []int {

	for i := 0; i < len(A); i++ {
		A[i] = A[i]*A[i]
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	return A
}

func main() {

}
