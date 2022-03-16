package _go

/*
剑指 Offer 41. 数据流中的中位数
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
*/
type MedianFinder struct {
	arr []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (m *MedianFinder) AddNum(num int) {
	if len(m.arr) == 0 {
		m.arr = append(m.arr, num)
		return
	}
	left := 0
	right := len(m.arr)
	var idx int
	for left < right {

		idx = (left + right) >> 1
		if m.arr[idx] < num {
			if left == idx {
				break
			}
			left = idx
		} else if m.arr[idx] > num {
			if right == idx {
				break
			}
			right = idx
		} else {
			break
		}

	}
	m.arr = append(m.arr, 0)
	if num > m.arr[idx] {
		copy(m.arr[idx+2:], m.arr[idx+1:])
		m.arr[idx+1] = num
	} else {
		copy(m.arr[idx+1:], m.arr[idx:])
		m.arr[idx] = num
	}

}

func (m *MedianFinder) FindMedian() float64 {
	if len(m.arr) == 0 {
		return 0
	}
	if (len(m.arr) % 2) == 0 {
		idx := len(m.arr) / 2
		return float64(m.arr[idx]+m.arr[idx-1]) / 2
	}

	return float64(m.arr[(len(m.arr) / 2)])

}
