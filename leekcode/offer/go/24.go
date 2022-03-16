package _go


/*
面试题24. 反转链表

定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL


限制：

0 <= 节点个数 <= 5000
*/

func reverseList(head *ListNode) *ListNode {
    return ReverseList(head)
}

func ReverseList(head *ListNode) *ListNode {
    var preNode *ListNode
    preNode = nil
    //  后一个节点
    nextNode := &ListNode{}
    nextNode = nil
    for head != nil {
        //  保存头节点的下一个节点，
        nextNode = head.Next
        //  将头节点指向前一个节点
        head.Next = preNode
        //  更新前一个节点
        preNode = head
        //  更新头节点
        head = nextNode
    }
    return preNode
}
