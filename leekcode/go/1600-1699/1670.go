package _600_1699

/*
1670. 设计前中后队列
请你设计一个队列，支持在前，中，后三个位置的 push 和 pop 操作。

请你完成 FrontMiddleBack 类：

FrontMiddleBack() 初始化队列。
void pushFront(int val) 将 val 添加到队列的 最前面 。
void pushMiddle(int val) 将 val 添加到队列的 正中间 。
void pushBack(int val) 将 val 添加到队里的 最后面 。
int popFront() 将 最前面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
int popMiddle() 将 正中间 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
int popBack() 将 最后面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
请注意当有 两个 中间位置的时候，选择靠前面的位置进行操作。比方说：

将 6 添加到 [1, 2, 3, 4, 5] 的中间位置，结果数组为 [1, 2, 6, 3, 4, 5] 。
从 [1, 2, 3, 4, 5, 6] 的中间位置弹出元素，返回 3 ，数组变为 [1, 2, 4, 5, 6] 。

*/

type FrontMiddleBackQueue struct {
	Queue []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (q *FrontMiddleBackQueue) PushFront(val int) {

	if cap(q.Queue) > len(q.Queue) { // 手动move
		q.Queue = append(q.Queue, 0)
		for i := len(q.Queue) - 1; i > 0; i-- {
			q.Queue[i] = q.Queue[i-1]
		}
		q.Queue[0] = val
		return
	}

	q.Queue = append([]int{val}, q.Queue...)

}

func (q *FrontMiddleBackQueue) PushMiddle(val int) {

	if q.IsEmpty() {
		q.PushBack(val)
		return
	}

	index := len(q.Queue) / 2

	if cap(q.Queue) > len(q.Queue) { // 手动move
		q.Queue = append(q.Queue, 0)
		for i := len(q.Queue) - 1; i > index; i-- {
			q.Queue[i] = q.Queue[i-1]
		}
		q.Queue[index] = val
		return
	}

	q.Queue = append(q.Queue[:index], append([]int{val}, q.Queue[index:]...)...)
}

func (q *FrontMiddleBackQueue) PushBack(val int) {
	q.Queue = append(q.Queue, val)

}

func (q *FrontMiddleBackQueue) PopFront() int {
	if q.IsEmpty() {
		return -1
	}

	val := q.Queue[0]
	for i := 1; i < len(q.Queue); i++ {
		q.Queue[i-1] = q.Queue[i]
	}
	q.Queue = q.Queue[:len(q.Queue)-1]

	return val
}

func (q *FrontMiddleBackQueue) PopMiddle() int {
	if q.IsEmpty() {
		return -1
	}
	if len(q.Queue) <= 2 {
		return q.PopFront()
	}

	index := len(q.Queue) / 2
	if len(q.Queue)%2 == 0 {
		index--
	}

	val := q.Queue[index]

	for i := index + 1; i < len(q.Queue); i++ {
		q.Queue[i-1] = q.Queue[i]
	}

	q.Queue = q.Queue[:len(q.Queue)-1]
	return val
}

func (q *FrontMiddleBackQueue) PopBack() int {
	if q.IsEmpty() {
		return -1
	}

	val := q.Queue[len(q.Queue)-1]
	q.Queue = q.Queue[:len(q.Queue)-1]
	return val
}

func (q *FrontMiddleBackQueue) IsEmpty() bool {
	return len(q.Queue) == 0
}

