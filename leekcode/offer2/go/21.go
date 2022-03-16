package _go


/*
剑指 Offer II 021. 删除链表的倒数第 n 个结点
给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。



示例 1：



输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	fast := head
	for n > 0 && fast != nil{
		fast = fast.Next
		n--
	}

	// 如果快指针为空,说明删除head节点
	if fast == nil {
		head = head.Next
		return head
	}

	// 双指针往后走
	cur := head
	for fast.Next != nil {
		cur = cur.Next
		fast = fast.Next
	}
	cur.Next = cur.Next.Next
	return head
}
