package btree


func min(n *node) Item {
    if n == nil {
        return nil
    }
    for len(n.children) > 0 {
        n = n.children[0]
    }
    if len(n.items) == 0 {
        return nil
    }
    return n.items[0]
}

func max(n *node) Item {
    if n == nil {
        return nil
    }
    for len(n.children) > 0 {
        n = n.children[len(n.children)-1]
    }
    if len(n.items) == 0 {
        return nil
    }
    return n.items[len(n.items)-1]
}


type Int int

// Less returns true if int(a) < int(b).
func (a Int) Less(b Item) bool {
    return a < b.(Int)
}