package _go


/*
面试题 02.04. 分割链表
编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
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