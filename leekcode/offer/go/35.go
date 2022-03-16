package _go

import "unsafe"

/*
剑指 Offer 35. 复杂链表的复制
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
*/

type Node struct {
    Val    int
    Next   *Node
    Random *Node
}

var table map[uintptr]*Node

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    table = make(map[uintptr]*Node)
    return head.Clone()
}

func (n *Node) Clone() *Node {
    var cur = &Node{
        Val: n.Val,
    }
    table[uintptr(unsafe.Pointer(n))] = cur
    if n.Next != nil {
        cur.Next = n.Next.Clone()
    }
    if n.Random != nil {
        cur.Random = table[uintptr(unsafe.Pointer(n.Random))]
    }

    return cur
}