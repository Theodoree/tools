package _00_499

import "fmt"

func findDisappearedNumbers(nums []int) []int {

	array := make([]int, len(nums)+1)

	for _, v := range nums {
		array[v]--
	}

	var array1 []int

	for i, v := range array {
		if v == 0 && i > 0 {
			array1 = append(array1, i)
		}
	}

	return array1
}

func main() {
	f := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(f)
}
