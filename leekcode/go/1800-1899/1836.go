package _800_1899



/*
1836. 从未排序的链表中移除重复元素
给定一个链表的第一个节点 head ，找到链表中所有出现多于一次的元素，并删除这些元素所在的节点。

返回删除后的链表。
*/
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	var arr [10e5]int

	cur:=head
	for cur != nil {
		arr[cur.Val]++
		cur = cur.Next
	}


	for head != nil && arr[head.Val] > 1 {
		head = head.Next
	}

	parent:=head
	cur=head
	for cur != nil {
		if arr[cur.Val] >1 {
			parent.Next = cur.Next
			cur = parent.Next
			continue
		}

		parent = cur
		cur = cur.Next
	}



	return head
}

