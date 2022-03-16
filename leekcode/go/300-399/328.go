package _00_399



/*
328. 奇偶链表

给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
*/

func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return head
    }
    odd := new(ListNode)
    event := new(ListNode)
    eventStart := head.Next // 这个等于偶数
    for i, node := 1, head; node != nil; i++ {
        if i%2 == 1 {
            odd.Next = node
            odd = node
        } else {
            event.Next = node
            event = node
        }
        node = node.Next
    }
    odd.Next = eventStart
    event.Next = nil
    return head
}


func oddEvenList(head *ListNode) *ListNode {

    if head == nil || head.Next == nil || head.Next.Next == nil {
        return head
    }

    var j, o *ListNode
    var jcur, ocur *ListNode
    var cnt int
    cnt = 1
    for head != nil {

        if cnt%2 != 0 {
            if jcur == nil {
                j, jcur = head, head
            } else {
                jcur.Next = head
                jcur = jcur.Next
            }
        } else {
            if ocur == nil {
                o, ocur = head, head
            } else {
                ocur.Next = head
                ocur = ocur.Next
            }
        }
        head = head.Next
        cnt++
    }

    if jcur != nil {
        jcur.Next = o
        if ocur != nil {
            ocur.Next = nil
        }

    }
    return j
}
