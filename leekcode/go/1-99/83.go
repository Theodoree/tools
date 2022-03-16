package __99


func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}


/*
func deleteDuplicates(head *ListNode) *ListNode {


	if head == nil{
		return  nil
	}

	current := head
	next := head.Next

	for next != nil {

		if current.Val == next.Val {
			next = next.Next
            current.Next = next
		}else{
			current.Next = next
			next = next.Next
			current = current.Next
		}

	}

	return  head
}
*/
