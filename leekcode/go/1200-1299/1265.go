package _200_1299

/*
1265. 逆序打印不可变链表
给您一个不可变的链表，使用下列接口逆序打印每个节点的值：

ImmutableListNode: 描述不可变链表的接口，链表的头节点已给出。
您需要使用以下函数来访问此链表（您 不能 直接访问 ImmutableListNode）：

ImmutableListNode.printValue()：打印当前节点的值。
ImmutableListNode.getNext()：返回下一个节点。
输入只用来内部初始化链表。您不可以通过修改链表解决问题。也就是说，您只能通过上述 API 来操作链表。
*/

func printLinkedListInReverse(head ImmutableListNode) {
	if head == nil {
		return
	}
	printLinkedListInReverse(head.getNext())
	head.printValue()
}
