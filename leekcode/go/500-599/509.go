package _00_599

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	nums := make([]int, N+1)
	nums[0] = 0
	nums[1] = 1
	for i := 2; i <= N; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums[N]
}

/*
func fib(N int) int {
	if N == 0{
		return 0
	}
	if N <= 2{
		return 1
	}


	return fib(N-1)+fib(N-2)
}
*/
