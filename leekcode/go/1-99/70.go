package __99

import "fmt"

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	first := 1
	second := 2

	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

func minSwap(nums1 []int, nums2 []int) int {

	for i:=0;i<len(nums1)-1;i++{
		if nums1[i] > nums1[i+1]{
			nums1[i]
		}else if nums2[i] > nums1[i+1]{

		}
	}

}
	fmt.Println(climbStairs(2))
}
