package __99

import (
	"fmt"
)

// 11. 盛最多水的容器 双向指针

func maxArea(height []int) int {
	res := 0
	i, j := 0, len(height)-1
	for i < j {
		res = max(res, (j-i) * min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))


}
