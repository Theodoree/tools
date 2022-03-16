package _00_299

//反转链表
func reverseList(head *ListNode) *ListNode {
    return ReverseList(head)
}

func ReverseList(head *ListNode) *ListNode {
    //  前一个节点
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