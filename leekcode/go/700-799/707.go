package _00_799


/*
707. 设计链表

设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。


示例：

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3


提示：

所有val值都在 [1, 1000] 之内。
操作次数将在  [1, 1000] 之内。
请不要使用内置的 LinkedList 库。
*/

type MyLinkedList struct {
    Head  *Node
    Tail  *Node
    total int
}

type Node struct {
    prev  *Node
    Next  *Node
    Value int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    n := &Node{
        prev:  nil,
        Next:  nil,
        Value: 0,
    }

    return MyLinkedList{Head: n, Tail: n}

}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if this.total-1 < index || index < 0{
        return -1
    }

    cur := this.Head.Next
    for index > 0 {
        cur = cur.Next
        index--
    }

    return cur.Value
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
    n := &Node{
        prev:  this.Head,
        Next:  this.Head.Next,
        Value: val,
    }
    this.Head.Next = n
    if n.Next != nil {
        n.Next.prev = n
    } else {
        this.Tail = n
    }
    this.total++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
    n := &Node{
        prev:  this.Tail,
        Next:  nil,
        Value: val,
    }
    n.prev.Next = n
    this.Tail = n
    this.total++

}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {

    switch {
    case index <= 0:
        this.AddAtHead(val)
        return
    case index == this.total:
        this.AddAtTail(val)
        return
    case this.total < index:
        return

    }

    cur := this.Head.Next

    for index > 0 {
        cur = cur.Next
        index--
    }

    parent := cur.prev
    next := cur

    n := &Node{
        prev:  parent,
        Next:  next,
        Value: val,
    }

    n.prev.Next = n
    if n.Next != nil {
        n.Next.prev = n
    }
    this.total++

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
    if this.total-1 < index || index < 0{
        return
    }

    cur := this.Head.Next
    for index > 0 {
        cur = cur.Next
        index--
    }

    if cur == this.Tail {
        parent := this.Tail.prev
        parent.Next = nil
        this.Tail = parent
    } else {
        parent := cur.prev
        next := cur.Next
        parent.Next = next
        if next != nil {
            next.prev = parent
        }
    }

    this.total--

}
