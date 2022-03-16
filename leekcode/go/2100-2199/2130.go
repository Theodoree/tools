package _100_2199


/*
2130. 链表最大孪生和
在一个大小为 n 且 n 为 偶数 的链表中，对于 0 <= i <= (n / 2) - 1 的 i ，第 i 个节点（下标从 0 开始）的孪生节点为第 (n-1-i) 个节点 。

比方说，n = 4 那么节点 0 是节点 3 的孪生节点，节点 1 是节点 2 的孪生节点。这是长度为 n = 4 的链表中所有的孪生节点。
孪生和 定义为一个节点和它孪生节点两者值之和。

给你一个长度为偶数的链表的头节点 head ，请你返回链表的 最大孪生和 。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}
func pairSum(head *ListNode) int {

	var buf = make([]int,0,100)

	cur:=head
	for cur != nil {
		buf = append(buf,cur.Val)
		cur = cur.Next
	}

	var max int
	for i:=0;i<len(buf)/2;i++{
		if buf[i] + buf[len(buf)-i-1] >= max {
			max = buf[i] + buf[len(buf)-i-1]
		}
	}



	return max


}

