package _00_199

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    filterMap := make(map[*ListNode]struct{})

    current := head
    for {
        if current == nil || current.Next == nil{
            return false
        }

        if _, ok := filterMap[current]; !ok {
            filterMap[current] = struct{}{}
        } else {
            return true
        }

        current = current.Next

    }
}
