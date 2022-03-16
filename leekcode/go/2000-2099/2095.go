package _000_2099


type ListNode struct {
	Val  int
	Next *ListNode
}
/*
2095. 删除链表的中间节点
给你一个链表的头节点 head 。删除 链表的 中间节点 ，并返回修改后的链表的头节点 head 。
长度为 n 链表的中间节点是从头数起第 ⌊n / 2⌋ 个节点（下标从 0 开始），其中 ⌊x⌋ 表示小于或等于 x 的最大整数。
对于 n = 1、2、3、4 和 5 的情况，中间节点的下标分别是 0、1、1、2 和 2 。
*/
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	parent:=head
	cur:=head
	next:=head
	for next != nil {
		if next.Next != nil {
			next = next.Next.Next
		}else{
			next = nil
			break
		}
		parent = cur
		cur = cur.Next
	}

	parent.Next = cur.Next
	return head
}

