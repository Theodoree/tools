package _go

/*
面试题 02.08. 环路检测
给定一个有环链表，实现一个算法返回环路的开头节点。
有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。
*/

func detectCycle(head *ListNode) *ListNode {

    m := make(map[*ListNode]struct{})

    cur := head
    for cur != nil {
        if _, ok := m[cur]; ok {
            return cur
        }
        m[cur] = struct{}{}
        cur = cur.Next
    }
    return nil
}