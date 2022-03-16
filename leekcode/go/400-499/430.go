package _00_499

import "runtime"

/*
430. 扁平化多级双向链表
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。
*/

type Node struct {
    Val   int
    Prev  *Node
    Next  *Node
    Child *Node
}

func flatten(root *Node) *Node {
    if root == nil {
        return nil
    }

    slice := root.GetSlice()

    for i := 1; i < len(slice); i++ {
        slice[i-1].Next = slice[i]
        slice[i].Prev = slice[i-1]
        slice[i-1].Child = nil
        slice[i].Child = nil

    }
	runtime.GC()
    return slice[0]
}

func (n *Node) GetSlice() []*Node {
    var arr []*Node
    cur := n
    for cur != nil {
        arr = append(arr, cur)
        if cur.Child != nil {
            arr = append(arr, cur.Child.GetSlice()...)
        }
        cur = cur.Next
    }
    return arr
}