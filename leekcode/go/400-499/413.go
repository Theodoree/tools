package _00_499


func numberOfArithmeticSlices(nums []int) (ans int) {
	n := len(nums)
	if n == 1 {
		return
	}

	d, t := nums[0]-nums[1], 0
	// 因为等差数列的长度至少为 3，所以可以从 i=2 开始枚举
	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == d {
			t++
		} else {
			d, t = nums[i-1]-nums[i], 0
		}
		ans += t
	}
	return
}

