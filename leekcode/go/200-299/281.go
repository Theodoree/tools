package _00_299

/*
281. 锯齿迭代器
给出两个一维的向量，请你实现一个迭代器，交替返回它们中间的元素。
*/
type ZigzagIterator struct {
	arr [2][]int
	count int
	flag bool

}

func Constructor(v1, v2 []int) *ZigzagIterator {
	return &ZigzagIterator{arr: [2][]int{v1,v2},count: len(v1)+len(v2),flag: true}
}

func (this *ZigzagIterator) next() int {
	this.count--


	var val int
	if this.flag && len(this.arr[0]) > 0 {
		val =this.arr[0][0]
		this.arr[0] = this.arr[0][1:]
		if len(this.arr[1]) > 0 {
			this.flag = false
		}
	}else if len(this.arr[1]) > 0 {
		val =this.arr[1][0]
		this.arr[1] = this.arr[1][1:]
		this.flag = true
	}

	return val
}

func (this *ZigzagIterator) hasNext() bool {
	return this.count > 0
}


