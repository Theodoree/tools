package __99



/*
86. 分隔链表

给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
*/
func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }

    head1 := &ListNode{Val: 0}
    head2 := &ListNode{Val: 0}
    cur1 := head1
    cur2 := head2

    for head != nil {
        if head.Val < x {
            cur1.Next = head
            head = head.Next
            cur1 = cur1.Next
            cur1.Next = nil
        } else {
            cur2.Next = head
            head = head.Next
            cur2 = cur2.Next
            cur2.Next = nil
        }

    }
    cur1.Next = head2.Next

    return head1.Next
}
