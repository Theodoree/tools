package _00_399


/*
369. 给单链表加一

用一个 非空 单链表来表示一个非负整数，然后将这个整数加一。

你可以假设这个整数除了 0 本身，没有任何前导的 0。

这个整数的各个数位按照 高位在链表头部、低位在链表尾部 的顺序排列。

示例:

输入: [1,2,3]
输出: [1,2,4]
*/

func plusOne(head *ListNode) *ListNode {

    if head == nil {
        return nil
    }

    var m = map[int]*ListNode{}

    cur := head
    cnt := 1
    for cur != nil {
        m[cnt] = cur
        cur = cur.Next
        cnt++
    }

    var max int
    max = len(m)
    for (max >= 1 && m[max].Val >= 10) || max == len(m) {

        if max == len(m) {
            m[max].Val += 1
        } else if max == 1 {
            // max = 1  head节点
            n := &ListNode{
                Val:  m[max].Val / 10,
                Next: nil,
            }
            m[max].Val %= 10
            n.Next = head
            head = n
            break
        }

        if m[max].Val >= 10 {
            if max-1 >= 1 {
                m[max-1].Val += m[max].Val / 10
            } else {
                n := &ListNode{
                    Val:  m[max].Val / 10,
                    Next: nil,
                }
                n.Next = head
                head = n
            }

            m[max].Val %= 10
        }

        max--
    }

    return head
}

