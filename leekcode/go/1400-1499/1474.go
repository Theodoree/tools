package _400_1499

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1474. 删除链表 M 个节点之后的 N 个节点
给定链表 head 和两个整数 m 和 n. 遍历该链表并按照如下方式删除节点:
开始时以头节点作为当前节点.
保留以当前节点开始的前 m 个节点.
删除接下来的 n 个节点.
重复步骤 2 和 3, 直到到达链表结尾.
在删除了指定结点之后, 返回修改过后的链表的头节点.
*/
func deleteNodes(head *ListNode, m int, n int) *ListNode {

	var cur, parent *ListNode
	cur = head

	for cur != nil {
		curM := m
		for curM > 0 && cur != nil {
			parent = cur
			cur = cur.Next
			curM--
		}

		if cur == nil { // 已经到end了
			break
		}

		curN := n
		for cur != nil && curN > 0 {
			cur = cur.Next
			curN--
		}
		parent.Next = cur
		cur = parent.Next
	}
	return head
}
