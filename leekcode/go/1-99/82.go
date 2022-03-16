package __99



/*
82. 删除排序链表中的重复元素 II

给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
*/

func deleteDuplicates(head *ListNode) *ListNode {

    if head == nil {
        return nil
    }

    var hashMap = make(map[int]int)
    cur := head
    for cur != nil {

        hashMap[cur.Val]++
        cur = cur.Next
    }

    for head != nil && hashMap[head.Val] > 1 {
        head = head.Next
    }

    if head == nil {
        return nil
    }

    parent := head
    cur = head.Next

    for cur != nil {
        for cur != nil && hashMap[cur.Val] > 1 {
            cur = cur.Next
        }

        parent.Next = cur
        parent = parent.Next
        if cur != nil {
            cur = cur.Next
        }
    }

    return head

}
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }

    pre := &ListNode{0,nil}
    pre.Next = head
    result := pre

    for head != nil && head.Next != nil{
        if head.Val == head.Next.Val{
            head = head.Next
            for head != nil && head.Next != nil{
                if head.Val != head.Next.Val {
                    break
                }
                head = head.Next
            }
            pre.Next = head.Next
            head = head.Next
        }else{
            pre = pre.Next
            head = head.Next
        }
    }

    return result.Next
}

