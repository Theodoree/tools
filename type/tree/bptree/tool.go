package bptree

func cut(length int) int {
    if length%2 == 0 {
        return length / 2
    }
    return length/2 + 1
}

