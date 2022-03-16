package _go

/*
剑指 Offer II 058. 日程表
请实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。
MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。
当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。
每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。否则，返回 false 并且不要将该日程安排添加到日历中。
请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)
*/
type MyCalendar struct {
	arr [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	if len(c.arr) == 0 {
		c.arr = append(c.arr, [2]int{start, end})
		return true
	}

	// 从头部插入
	if c.arr[0][0] >= end {
		if c.arr[0][0] == end {
			c.arr[0][0] = start
		} else {
			c.arr = append(c.arr, [2]int{})
			copy(c.arr[1:], c.arr)
			c.arr[0][0] = start
			c.arr[0][1] = end
		}
		return true
	}

	if c.arr[len(c.arr)-1][1] <= start {
		if c.arr[len(c.arr)-1][1] == start {
			c.arr[len(c.arr)-1][1] = end
		} else {
			c.arr = append(c.arr, [2]int{start, end})
		}
		return true
	}

	// 中间插入,二分查找
	left := 0
	right := len(c.arr)
	mid := 0
	for left < right {
		mid = (left + right) / 2
		if c.arr[mid][1] < end {
			left++
		} else if c.arr[mid][1] > end {
			if mid == right {
				break
			}
			right = mid
		} else {
			return false // 等于end,直接返回就行了
		}
	}

	switch {
	case c.arr[mid][1] > start:
		return false
	case c.arr[mid+1][0] < start && c.arr[mid+1][1] > start:
		return false
	case c.arr[mid+1][0] < end && c.arr[mid+1][1] > end:
		return false
	case c.arr[mid+1][0] == end:
		c.arr[mid+1][0] = start
		return true
	case c.arr[mid+1][1] == start:
		c.arr[mid+1][1] = end
		return true
	}

	c.arr = append(c.arr, [2]int{})
	copy(c.arr[left+1:], c.arr[left:])
	c.arr[left][0] = start
	c.arr[left][1] = end

	return true
}
