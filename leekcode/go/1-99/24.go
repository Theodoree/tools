package __99

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil{
		return  head
	}

	next:=head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return  next
}
