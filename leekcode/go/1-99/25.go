package __99


/*
25. K 个一组翻转链表
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。



示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	var m []*ListNode
	for head != nil {
		next := head.Next
		head.Next = nil
		m = append(m, head)
		head = next
	}

	head = m[k-1]
	var lastTail *ListNode

	for len(m) > 0 {
		for i := k-1; i > 0; i-- {
			m[i].Next = m[i-1]
		}
		lastTail = m[0]
		m = m[k:]

		if len(m) < k {
			for _, v := range m {
				lastTail.Next = v
				lastTail = v
			}
			break
		}

		lastTail.Next = m[k-1]
	}

	return head

}
