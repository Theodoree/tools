package _go

// 移除重复节点

func removeDuplicateNodes(head *ListNode) *ListNode {
    m := make(map[int]struct{})
    var cur *ListNode
    var n *ListNode
    for head != nil {
        if _, ok := m[head.Val]; ok {
            head = head.Next
            continue
        }

        m[head.Val] = struct {}{}

        k := *head
        k.Next = nil

        if cur == nil {
            cur = &k
            n = &k
        } else {
            cur.Next = &k
            cur = cur.Next
        }
        head = head.Next
    }

    return n
}
