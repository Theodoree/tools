package _00_399


/*
346. 数据流中的移动平均值

给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算其所有整数的移动平均值。

示例:

MovingAverage m = new MovingAverage(3);
m.next(1) = 1
m.next(10) = (1 + 10) / 2
m.next(3) = (1 + 10 + 3) / 3
m.next(5) = (10 + 3 + 5) / 3
*/
type MovingAverage struct {
    array []int
    cnt   int
    last  int
    Index int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    return MovingAverage{array: make([]int, size)}
}

func (this *MovingAverage) Next(val int) float64 {
    if this.cnt < len(this.array) {
        this.array[this.cnt] = val
        this.last += val
        this.cnt++
        this.Index++
    } else {
        if this.Index == this.cnt {
            this.Index = 0
        }
        this.last += val - this.array[this.Index]
        this.array[this.Index] = val
        this.Index++
    }

    return float64(this.last) / float64(this.cnt)
}