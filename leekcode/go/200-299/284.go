package _00_299

/*
284. 窥探迭代器
请你在设计一个迭代器，在集成现有迭代器拥有的 hasNext 和 next 操作的基础上，还额外支持 peek 操作。

实现 PeekingIterator 类：

PeekingIterator(Iterator<int> nums) 使用指定整数迭代器 nums 初始化迭代器。
int next() 返回数组中的下一个元素，并将指针移动到下个元素处。
bool hasNext() 如果数组中存在下一个元素，返回 true ；否则，返回 false 。
int peek() 返回数组中的下一个元素，但 不 移动指针。
注意：每种语言可能有不同的构造函数和迭代器，但均支持 int next() 和 boolean hasNext() 函数。
*/

type PeekingIterator struct {
	itre *Iterator

	curVal int
}

func Constructor(iter *Iterator) *PeekingIterator {
	p := &PeekingIterator{
		itre: iter,
	}

	if iter.hasNext() {
		p.curVal = iter.next()
	}

	return p
}

func (this *PeekingIterator) hasNext() bool {
	return this.curVal != 0
}

func (this *PeekingIterator) next() int {
	c := this.curVal
	this.curVal = 0
	if this.itre.hasNext() {
		this.curVal = this.itre.next()
	}
	return c
}

func (this *PeekingIterator) peek() int {
	return this.curVal
}
