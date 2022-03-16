package _00_799


/*
708. 循环有序列表的插入
给定循环单调非递减列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环非降序的。
给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。
如果有多个满足条件的插入位置，你可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。
如果列表为空（给定的节点是 null），你需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。
*/
func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		aNode := &Node{Val: x}
		aNode.Next = aNode
		return aNode
	}

	curr := aNode
	for !(curr.Val <= x && x <= curr.Next.Val || //升序中间 直接插入
		curr.Next.Val < curr.Val && (curr.Val <= x || x <= curr.Next.Val) || // curr < prev 时 必定 prev = tail & curr = head  ,如果正好可以插入 就插入
		curr.Next == aNode) { // 相等时，上述条件进入死循环，所以判断一下
		curr = curr.Next
	}
	curr.Next = &Node{Val: x, Next: curr.Next}

	return aNode
}

