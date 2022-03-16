package _00_299

import "fmt"

//217.存在重复元素

func main() {
	f := []int{1, 2, 3, 4}
	ok := containsDuplicate(f)
	fmt.Println(ok)
}

func containsDuplicate(nums []int) bool {
	v := make(map[int]int)

	for _, num := range nums {
		if _, ok := v[num]; ok {
			return true
		} else {
			v[num] = 1
		}
	}

	return false
}
