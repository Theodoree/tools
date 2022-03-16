package _000_2099

/*
2058. 找出临界点之间的最小和最大距离
链表中的 临界点 定义为一个 局部极大值点 或 局部极小值点 。
如果当前节点的值 严格大于 前一个节点和后一个节点，那么这个节点就是一个  局部极大值点 。
如果当前节点的值 严格小于 前一个节点和后一个节点，那么这个节点就是一个  局部极小值点 。
注意：节点只有在同时存在前一个节点和后一个节点的情况下，才能成为一个 局部极大值点 / 极小值点 。
给你一个链表 head ，返回一个长度为 2 的数组 [minDistance, maxDistance] ，其中 minDistance 是任意两个不同临界点之间的最小距离，maxDistance 是任意两个不同临界点之间的最大距离。如果临界点少于两个，则返回 [-1，-1] 。

*/
func nodesBetweenCriticalPoints(head *ListNode) []int {

	min := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}

	startNode := -1
	lastNode := -1

	minResult := math.MaxInt64

	var next *ListNode
	idx := 0
	parent := head
	cur := head.Next
	if cur != nil {
		next = cur.Next
	}
	for next != nil {
		if (cur.Val > parent.Val && cur.Val > next.Val) || (cur.Val < parent.Val && cur.Val < next.Val) {
			if startNode == -1 {
				startNode = idx
			}
			if lastNode != -1 {
				minResult = min(minResult, idx-lastNode)
			}
			lastNode = idx
		}

		parent = cur
		cur = next
		next = next.Next
		idx++
	}

	if minResult == math.MaxInt64 {
		return []int{-1, -1}
	}

	return []int{minResult, lastNode - startNode}
}
func init() {
	debug.SetGCPercent(0)
}
