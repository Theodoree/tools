package __99

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 19. 删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 0 {
		return nil
	}

	current := head
	NSetpNode := current
	for i := 0; i < n; i++ {
		NSetpNode = NSetpNode.Next
	}

	if NSetpNode == nil {
		return head.Next
	}

	for {

		if NSetpNode.Next == nil {
			current.Next = current.Next.Next
			break
		} else if NSetpNode.Next != nil {
			NSetpNode = NSetpNode.Next
			current = current.Next
		}
	}

	return head
}

func main() {
	fmt.Println(removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 2))
	//fmt.Println(removeNthFromEnd(&ListNode{Val:1},1))

}
