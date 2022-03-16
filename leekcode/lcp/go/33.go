package _go


func storeWater(bucket []int, vat []int) int {

	min := func(a, b int)int {
		if a>b {return b}
		return a
	}

	finish := true
	for _, v := range vat {
		if v != 0 {
			finish = false
			break
		}
	}
	if finish {return 0}

	//需要额外补多少下才能满足能在time次内完成
	ans := int(2*1e4)

	for time := 1; time<=int(1e4); time++ {
		finish := true
		add := 0
		for i, _ := range bucket {
			n := int(math.Ceil(float64(vat[i])/float64(time)))
			if n > bucket[i] {
				add += n-bucket[i]
				finish = false
			}
		}

		if finish {
			return min(ans, time)
		}

		ans = min(time+add, ans)
	}

	return ans
}
