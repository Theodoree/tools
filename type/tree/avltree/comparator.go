package avltree


// a < b -1
// a == b 0
// a > b  1
type Comparator func(a,b interface{}) int


func stringComparator(a, b interface{}) int {
    s1 := a.(string)
    s2 := b.(string)
    min := len(s2)
    if len(s1) < len(s2) {
        min = len(s1)
    }
    diff := 0
    for i := 0; i < min && diff == 0; i++ {
        diff = int(s1[i] - s2[i])
    }
    if diff == 0 {
        diff = len(s1) - len(s2)
    }
    if diff < 0 {
        return -1
    }
    if diff > 0 {
        return 1
    }
    return 0
}


func intComparator(a, b interface{}) int {
    aInt := a.(int)
    bInt := b.(int)
    switch {
    case aInt > bInt:
        return 1
    case aInt < bInt:
        return -1
    default:
        return 0
    }
}
