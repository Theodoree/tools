package _00_299


/*
203. 移除链表元素

删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/
type ListNode struct {
    Val  int
    Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
    pre := ListNode{0, head}
    tmp := &pre
    for tmp.Next != nil { //迭代
        if tmp.Next.Val == val {
            tmp.Next = tmp.Next.Next
        } else {
            tmp = tmp.Next
        }
    }
    return pre.Next
}

func removeElements(head *ListNode, val int) *ListNode {

    if head == nil {
        return nil
    }

    for head != nil&& head.Val == val {
        head = head.Next
    }
    if head != nil {
        head.Next = removeElements(head.Next, val) //递归
    }
    return head
}
