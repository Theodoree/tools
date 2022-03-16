package __99

/*
23. 合并K个排序链表
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度

分治大法好
*/


type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 != nil {
		return l2
	}

	if l1 == nil || (l2 != nil && l1.Val > l2.Val) {
		l1, l2 = l2, l1
	}
	if l1 != nil {
		l1.Next = MergeTwoLists(l1.Next, l2)
	}
	return l1
}

func mergeKLists(lists []*ListNode) *ListNode {
	listsLen := len(lists)
	switch listsLen {
	case 0:
		return nil
	case 1:
		return lists[0]
	default:
		return MergeTwoLists(mergeKLists(lists[:listsLen/2]), mergeKLists(lists[listsLen/2:]))
	}
}
func main(){



}