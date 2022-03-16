package _go



/*
剑指 Offer II 041. 滑动窗口的平均值
给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算滑动窗口里所有数字的平均值。

实现 MovingAverage 类：

MovingAverage(int size) 用窗口大小 size 初始化对象。
double next(int val) 成员函数 next 每次调用的时候都会往滑动窗口增加一个整数，请计算并返回数据流中最后 size 个值的移动平均值，即滑动窗口里所有数字的平均值。


示例：

输入：
inputs = ["MovingAverage", "next", "next", "next", "next"]
inputs = [[3], [1], [10], [3], [5]]
输出：
[null, 1.0, 5.5, 4.66667, 6.0]

解释：
MovingAverage movingAverage = new MovingAverage(3);
movingAverage.next(1); // 返回 1.0 = 1 / 1
movingAverage.next(10); // 返回 5.5 = (1 + 10) / 2
movingAverage.next(3); // 返回 4.66667 = (1 + 10 + 3) / 3
movingAverage.next(5); // 返回 6.0 = (10 + 3 + 5) / 3


提示：

1 <= size <= 1000
-105 <= val <= 105
最多调用 next 方法 104 次
*/

type MovingAverage struct {
	arr   []int
	cnt   int
	index int
	total int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{arr: make([]int, size)}
}

func (m *MovingAverage) Next(val int) float64 {
	if m.cnt < len(m.arr) {
		m.arr[m.cnt] = val
		m.total += val
		m.cnt++
		m.index++
	} else {
		if m.index == m.cnt {
			m.index = 0
		}
		m.total += val - m.arr[m.index]
		m.arr[m.index] = val
		m.index++
	}

	return float64(m.total) / float64(m.cnt)
}