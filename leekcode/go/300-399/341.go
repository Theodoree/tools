package _00_399

/*
341. 扁平化嵌套列表迭代器
给你一个嵌套的整数列表 nestedList 。每个元素要么是一个整数，要么是一个列表；该列表的元素也可能是整数或者是其他列表。请你实现一个迭代器将其扁平化，使之能够遍历这个列表中的所有整数。
实现扁平迭代器类 NestedIterator ：
NestedIterator(List<NestedInteger> nestedList) 用嵌套列表 nestedList 初始化迭代器。
int next() 返回嵌套列表的下一个整数。
boolean hasNext() 如果仍然存在待迭代的整数，返回 true ；否则，返回 false 。
你的代码将会用下述伪代码检测：
*/
type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool           {}
func (this NestedInteger) GetInteger() int           {}
func (n *NestedInteger) SetInteger(value int)        {}
func (this *NestedInteger) Add(elem NestedInteger)   {}
func (this NestedInteger) GetList() []*NestedInteger {}

type NestedIterator struct {
	arr []int // 还有一个优化点就是不递归取出来,然后记住数组的下标
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var arr []int

	for _, v := range nestedList {
		bfs(v, &arr)
	}

	return &NestedIterator{arr}
}

func (this *NestedIterator) Next() int {
	val := this.arr[0]
	this.arr = this.arr[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	return len(this.arr) > 0
}

func bfs(n *NestedInteger, arr *[]int) {
	if n == nil {
		return
	}
	if n.IsInteger() {
		*arr = append(*arr, n.GetInteger())
		return
	}
	for _, v := range n.GetList() {
		bfs(v, arr)
	}
}
