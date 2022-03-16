package _00_799

/*
702. 搜索长度未知的有序数组

给定一个升序整数数组，写一个函数搜索 nums 中数字 target。如果 target 存在，返回它的下标，否则返回 -1。注意，这个数组的大小是未知的。你只可以通过 ArrayReader 接口访问这个数组，ArrayReader.get(k) 返回数组中第 k 个元素（下标从 0 开始）。

你可以认为数组中所有的整数都小于 10000。如果你访问数组越界，ArrayReader.get 会返回 2147483647。



样例 1：

输入: array = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 存在在 nums 中，下标为 4
样例 2：

输入: array = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不在数组中所以返回 -1

*/
type ArrayReader struct {
	v []int
}

func (a *ArrayReader) get(key int) int {
	if key >= len(a.v) {
		return 2147483647
	}

	return a.v[key]

}

//这个只用于少数数组
func search(reader ArrayReader, target int) int {
	var left, right int
	//查看数组有多长
	var cnt int
	for i := 1; i < math.MaxInt32; i++ {
		cnt++	//这里的渐近测试长度会耗费大量时间
		if reader.get(i*10) == 2147483647 {
			right = i * 10

			for reader.get(right+1) != 2147483647 {
				cnt++
				right--
			}
			break
		}
	}

	for left <= right {
		cnt++
		mid := (left + right) >> 1

		if reader.get(mid) > target {
			right = mid - 1
		} else if reader.get(mid) < target {
			left = mid + 1
		} else {
			return mid
		}

	}

	fmt.Println(cnt)

	return -1
}


func search1(reader ArrayReader, target int) int {
	left, right := 0, 2147483647
	var cnt int
	for left <= right {
		cnt++
		mid := (left + right) >> 1
		val := reader.get(mid)
		if val == target {
			return mid
		} else if val > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println(cnt)

	return -1
}

