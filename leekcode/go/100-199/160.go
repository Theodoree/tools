package _00_199


/*
160. 相交链表

编写一个程序，找到两个单链表相交的起始节点。
*/

type ListNode struct {
    Val  int
    Next *ListNode
}


func getIntersectionNode(headA, headB *ListNode) *ListNode {

    for headA.Next != nil {
        headA = headA.Next
    }

    for headB != nil {
    }

    filterMap := make(map[*ListNode]struct{})
    for headA != nil || headB != nil {
        if headA != nil {
            if _, ok := filterMap[headA]; ok {
                return headA
            } else {
                filterMap[headA] = struct{}{}
            }
            headA = headA.Next
        }
        if headB != nil {
            if _, ok := filterMap[headB]; ok {
                return headB
            } else {
                filterMap[headB] = struct{}{}
            }
            headB = headB.Next

        }
    }

    return nil

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil{
        return nil
    }
    first := headA
    second := headB
    for first != second{
        // 两个都为null或者找到起始节点就会跳出循环
        if first == nil{
            first = headB
        }else{
            first = first.Next
        }

        if second == nil{
            second = headA
        }else{
            second = second.Next
        }
    }
    return first

}