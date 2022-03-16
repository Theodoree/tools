package _00_199

import "unsafe"

/*
138. 复制带随机指针的链表
给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。

要求返回这个链表的 深拷贝。

我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

val：一个表示 Node.val 的整数。
random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
	Left   *Node
	Right  *Node
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
