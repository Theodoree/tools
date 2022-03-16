package __99



/*
92. 反转链表 II

反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {

    switch {
    case head == nil:
        return nil
    case m == n:
        return head
    }

    var cnt int
    cnt = 1
    cur := head

    var subHead, subTail, subTailNext *ListNode

    if m == 1 {
        subHead = head
    }

    for cur != nil {
        if cnt == m-1 && subHead == nil {
            subHead = cur
        }

        if cnt == n {
            subTail = cur
            break
        }
        cur = cur.Next
        cnt++
    }

    if subHead == nil {
        return nil
    }

    if subTail != nil {
        subTailNext = subTail.Next
        subTail.Next = nil
    }

    if m == 1 {
        head = ReverseList(head)
        cur = head
        if subTailNext != nil {

            for cur.Next != nil {
                cur = cur.Next
            }
            cur.Next = subTailNext
        }
    } else {
        subHead.Next = ReverseList(subHead.Next)

        cur = subHead.Next
        if subTailNext != nil {
            for cur.Next != nil {
                cur = cur.Next
            }
            cur.Next = subTailNext
        }
    }

    return head

}

func ReverseList(head *ListNode) *ListNode {
    //  前一个节点
    var preNode *ListNode
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
