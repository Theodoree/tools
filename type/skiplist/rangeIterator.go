package skiplist

type rangeIterator struct {
    iter
    upperLimit interface{}
    lowerLimit interface{}
}

// 重写方法
func (i *rangeIterator) Next() bool {
    if !i.current.hasNext() {
        return false
    }

    next := i.current.next()

    if !i.list.lessThan(next.key, i.upperLimit) { // 在范围内,当前节点往下扫描
        return false
    }

    i.current = i.current.next()
    i.key = i.current.key
    i.value = i.current.value
    return true
}

func (i *rangeIterator) Previous() bool {
    if !i.current.hasPrevious() {
        return false
    }

    previous := i.current.previous()

    if i.list.lessThan(previous.key, i.lowerLimit) { // 在范围内,当前节点往下扫描
        return false
    }

    i.current = i.current.previous()
    i.key = i.current.key
    i.value = i.current.value
    return true
}

func (i *rangeIterator) Seek(key interface{}) (ok bool) {
    if i.list.lessThan(key, i.lowerLimit) { // 检查key是否符合条件
        return
    } else if !i.list.lessThan(key, i.upperLimit) {
        return
    }

    return i.iter.Seek(key)
}

func (i *rangeIterator) Close() {
    i.iter.Close()
    i.upperLimit = nil
    i.lowerLimit = nil
}
