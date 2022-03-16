package _400_1499


/*
1429. 第一个唯一数字
给定一系列整数，插入一个队列中，找出队列中第一个唯一整数。

实现 FirstUnique 类：

FirstUnique(int[] nums) 用数组里的数字初始化队列。
int showFirstUnique() 返回队列中的 第一个唯一 整数的值。如果没有唯一整数，返回 -1。（译者注：此方法不移除队列中的任何元素）
void add(int value) 将 value 插入队列中。
*/

type FirstUnique struct {
	hashTable map[int]int
	queue    []int
	idx      int
}


func Constructor(nums []int) FirstUnique {
	var f FirstUnique
	f.hashTable = make(map[int]int)
	for _,v:=range nums{
		f.hashTable[v]++
	}
	f.queue = nums
	return f
}


func (f *FirstUnique) ShowFirstUnique() int {
	for f.idx < len(f.queue){
		if f.hashTable[f.queue[f.idx]] == 1 {
			return f.queue[f.idx]
		}
		f.idx++
	}
	return -1
}


func (f *FirstUnique) Add(value int)  {
	val:=f.hashTable[value]
	switch val {
	case 0:
		f.queue = append(f.queue,value)
		f.hashTable[value]++
	case 1:
		f.hashTable[value]++
	}
}
