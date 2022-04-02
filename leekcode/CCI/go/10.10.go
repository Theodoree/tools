package _go



/*
面试题 10.10. 数字流的秩
假设你正在读取一串整数。每隔一段时间，你希望能找出数字 x 的秩(小于或等于 x 的值的个数)。请实现数据结构和算法来支持这些操作，也就是说：
实现 track(int x) 方法，每读入一个数字都会调用该方法；
实现 getRankOfNumber(int x) 方法，返回小于或等于 x 的值的个数。
注意：本题相对原题稍作改动
*/
type StreamRank struct {
	counts [50002]int
}

func Constructor() StreamRank {
	return StreamRank{}
}

func (this *StreamRank) Track(x int) {
	for i := x + 1; i < 50002; i += i & -i {
		this.counts[i]++
	}
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	rank := 0
	for i := x + 1; i > 0; i -= i & -i {
		rank += this.counts[i]
	}
	return rank
}


