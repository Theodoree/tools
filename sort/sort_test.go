package sort

import (
	"golang.org/x/exp/constraints"
	"math/rand"
	"testing"
)

func AreSorted[T constraints.Integer](arr []T) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func randArr[T constraints.Integer](t T, cnt int) []T {
	var arr []T
	for i := 0; i < cnt; i++ {
		arr = append(arr, T(rand.Int()))
	}
	return arr
}

func TestBinarySearch(t *testing.T) {

	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("BinarySearchBigger", func(t *testing.T) {
		if BinarySearchBigger(arr, 5) != 6 {
			t.Fatal(arr)
		}
	})

	t.Run("BinarySearch", func(t *testing.T) {
		if BinarySearch(arr, 5) != 5 {
			t.Fatal(arr)
		}
	})

	t.Run("BinarySearchLast", func(t *testing.T) {
		var arr = []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10}

		if BinarySearchLast(arr, 5) != 6 {
			t.Fatal(arr)
		}
	})

	t.Run("BinarySearchLast", func(t *testing.T) {
		if BinarySearchLess(arr, 5) != 4 {
			t.Fatal(arr)
		}
	})

}

func TestQuickSort(t *testing.T) {

	t.Run("QuickSort", func(t *testing.T) {
		arr := randArr(0, 20)
		arr = QuickSort(arr)
		if !AreSorted(arr) {
			t.Fatal(arr)
		}
	})

	t.Run("QuickSort1", func(t *testing.T) {
		arr := randArr(0, 20)
		arr = QuickSort1(arr)
		if !AreSorted(arr) {
			t.Fatal(arr)
		}
	})

	t.Run("QuickSort2", func(t *testing.T) {
		arr := randArr(0, 20)
		arr = QuickSort2(arr)
		if !AreSorted(arr) {
			t.Fatal(arr)
		}
	})

	t.Run("QuickSort3", func(t *testing.T) {
		arr := randArr(0, 20)
		arr = QuickSort3(arr)
		if !AreSorted(arr) {
			t.Fatal(arr)
		}
	})

}

func TestMergeSort(t *testing.T) {
	t.Run("MergeSort", func(t *testing.T) {
		arr := randArr(0, 20)
		arr = MergeSort(arr)
		if !AreSorted(arr) {
			t.Fatal(arr)
		}

	})

	t.Run("MergeSortByDownToUp", func(t *testing.T) {
		arr := randArr(0, 20)
		arr = MergeSortByDownToUp(arr)
		if !AreSorted(arr) {
			t.Fatal(arr)
		}
	})
}

func TestShellSort(t *testing.T) {
	arr := randArr(0, 20)
	ShellSort(arr)
	if !AreSorted(arr) {
		t.Fatal(arr)
	}
}

func TestSelectSort(t *testing.T) {
	arr := randArr(0, 20)
	SelectSort(arr)
	if !AreSorted(arr) {
		t.Fatal(arr)
	}
}

func TestBubbleSort(t *testing.T) {
	arr := randArr(0, 20)
	BubbleSort(arr)
	if !AreSorted(arr) {
		t.Fatal(arr)
	}
}
