package _000_2099


/*
2046. 给按照绝对值排序的链表排序
给你一个链表的头结点 head ，这个链表是根据结点的绝对值进行升序排序, 返回重新根据节点的值进行升序排序的链表
*/
func sortLinkedList(head *ListNode) *ListNode {


	var sentinel ListNode
	var tail *ListNode
	for head != nil && head.Val < 0 {
		oldNext:=sentinel.Next
		sentinel.Next = head
		head = head.Next
		sentinel.Next.Next = oldNext
		if tail == nil {
			tail = sentinel.Next
		}
	}


	parent:=head
	cur:=head

	for cur != nil{

		if cur.Val < 0 {
			oldNext:=sentinel.Next
			n:=cur.Next
			sentinel.Next = cur
			sentinel.Next.Next = oldNext

			parent.Next = n // skip
			cur = n
			if tail == nil {
				tail = sentinel.Next
			}
		}else{
			parent = cur
			cur = cur.Next
		}
	}
	if tail == nil {
		return head
	}
	tail.Next = head




	return sentinel.Next
}
