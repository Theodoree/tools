package _800_1899

/*
1823. 找出游戏的获胜者
共有 n 名小伙伴一起做游戏。小伙伴们围成一圈，按 顺时针顺序 从 1 到 n 编号。确切地说，从第 i 名小伙伴顺时针移动一位会到达第 (i+1) 名小伙伴的位置，其中 1 <= i < n ，从第 n 名小伙伴顺时针移动一位会回到第 1 名小伙伴的位置。
游戏遵循如下规则：
从第 1 名小伙伴所在位置 开始 。
沿着顺时针方向数 k 名小伙伴，计数时需要 包含 起始时的那位小伙伴。逐个绕圈进行计数，一些小伙伴可能会被数过不止一次。
你数到的最后一名小伙伴需要离开圈子，并视作输掉游戏。
如果圈子中仍然有不止一名小伙伴，从刚刚输掉的小伙伴的 顺时针下一位 小伙伴 开始，回到步骤 2 继续执行。
否则，圈子中最后一名小伙伴赢得游戏。
给你参与游戏的小伙伴总数 n ，和一个整数 k ，返回游戏的获胜者。
*/

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func Create(arr []int) (*ListNode1, *ListNode1) {
	var sentry = &ListNode1{}

	var cur *ListNode1
	cur = sentry

	for _, v := range arr {
		cur.Next = &ListNode1{Val: v}
		cur = cur.Next
	}
	return sentry.Next, cur
}

func findTheWinner(n int, k int) int {
	var arr = make([]int, 0, n)
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	root, tail := Create(arr)
	tail.Next = root
	total := n

	parent := root
	cur := root

	for total != 1 {
		if root.Next == root {
			break
		}
		curK := k

		for i := 0; i < curK-1; i++ {
			parent = cur
			cur = cur.Next
		}
		parent.Next = cur.Next
		cur = parent.Next
		root = root.Next
		total--
	}

	return root.Val
}
