package _go


/*
面试题52. 两个链表的第一个公共节点
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := make(map[*ListNode]struct{})
    for headA != nil {
        m[headA] = struct{}{}
        headA = headA.Next
    }

    for headB != nil {
        if _, ok := m[headB]; ok {
            return headB
        }
        headB = headB.Next
    }

    return  nil

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    na, nb := headA, headB
    la, lb := 0, 0
    for nil != na {
        na = na.Next
        la++
    }
    for nil != nb {
        nb = nb.Next
        lb++
    }
    na, nb = headA, headB
    if la >= lb {
        d := la - lb
        for i := 0; i < d; i++ {
            na = na.Next
        }
    } else {
        d := lb - la
        for i := 0; i < d; i++ {
            nb = nb.Next
        }
    }
    for nil != na {
        if na == nb {
            return na
        }
        na, nb = na.Next, nb.Next
    }
    return nil
}
