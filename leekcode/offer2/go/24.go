package _go
/*
剑指 Offer II 024. 反转链表
给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。



示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：


输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]


提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var next *ListNode
	var cur *ListNode
	for head != nil {
		cur = head
		n:=head.Next
		cur.Next = next
		head =n
		next = cur
	}

	return cur
}