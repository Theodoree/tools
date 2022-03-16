package _00_299


//234. 回文链表

func isPalindrome(head *ListNode) bool {

    if head == nil || head.Next == nil {
        return true
    }

    //快慢指针
    last, slow, fast := head, head, head

    var prev,next *ListNode

    for fast != nil && fast.Next != nil {
        last = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    //pre为前半部分链表的最后节点
    last.Next = nil
    //反转后半部分链表

    for slow!=nil{
        next = slow.Next
        slow.Next = prev
        //移动prev slow
        prev =slow
        slow =next
    }

    //比较前后两个链表
    before,after:=head,prev
    for before!=nil{
        if before.Val!=after.Val{
            return false
        }
        before = before.Next
        after = after.Next
    }

    return true
}

