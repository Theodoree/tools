package _go

import "fmt"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var cnt int
    a, b := headA, headB
    for a != b {
        if a == nil {
            a = headB
        } else {
            a = a.Next
        }
        if b == nil {
            b = headA
        } else {
            b = b.Next
        }
        cnt++
    }
    fmt.Println(cnt)
    return a
}
