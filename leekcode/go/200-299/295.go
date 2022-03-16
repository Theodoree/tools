package _00_299

/*
295. 数据流的中位数
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
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
