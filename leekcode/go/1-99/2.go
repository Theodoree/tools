package __99

/*
2. 两数相加

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	Head := &ListNode{}
	carry := 0
	current := Head
	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
            l1 = l1.Next
            
        }
		if l2 != nil {
			val2 = l2.Val
            l2 = l2.Next
        }

		sum := carry + val1 + val2
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return Head.Next
}