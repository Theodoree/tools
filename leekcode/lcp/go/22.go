package _go

func paintingPlan(n int, k int) int {
	if n*n == k {
		return 1
	}
	var ans int
	var arr [][]int = [][]int{}
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if i*n+j*n-i*j == k {
				temp := []int{i, j}
				arr = append(arr, temp)
				break
			}
		}
	}
	for _, a := range arr {
		ans += Check(n, a[0]) * Check(n, a[1])
	}
	return ans
}

func Check(n, k int) int {
	var fz int = 1
	var fm int = 1
	for i := 1; i <= k; i++ {
		fz *= (n - i + 1)
	}
	for i := k; i >= 1; i-- {
		fm *= i
	}
	return fz / fm
}
