package __99



// 21.合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := ListNode{}
	c := &ret

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			c.Next = l1
			l1 = l1.Next
		} else {
			c.Next = l2
			l2 = l2.Next
		}
		c = c.Next
	}

	if l1 == nil {
		c.Next = l2
	} else if l2 == nil {
		c.Next = l1
	}

	return ret.Next
}
