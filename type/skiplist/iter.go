package skiplist

type iter struct {
    key, value interface{} // key,value
    current    *node       // 当前node
    list       *SkipList   // 跳过列表
}

func (i *iter) Key() interface{} {
    return i.key
}

func (i *iter) Value() interface{} {
    return i.value
}

// 迭代器 向后
func (i *iter) Next() bool {
    if !i.current.hasNext() { // 如果没有next
        return false
    }
    i.current = i.current.next()
    i.key = i.current.key
    i.value = i.current.value

    return true
}

// 迭代器 向前
func (i *iter) Previous() bool {
    if !i.current.hasPrevious() {
        return false
    }

    i.current = i.current.previous()
    i.key = i.current.key
    i.value = i.current.value
    return true
}

func (i *iter) Seek(key interface{}) bool {
    current := i.current
    list := i.list

    if current == nil { // 如果是初始化,那么current设置为list.header
        current = list.header
    }

    // 如果当前key大于key那么重置
    if current.key != key && list.lessThan(key, current.key) {
        current = list.header
    }

    // 如果当前前驱节点为空,那么重置current,否则从前驱开始扫描
    if current.backward == nil {
        current = list.header
    } else {
        current = current.backward
    }

    current = list.getPath(current, nil, key) // 寻找插入位置

    if current == nil { // 如果为空,说明该Key不存在
        return false
    }

    i.current = current
    i.key = current.key
    i.value = current.value

    return true
}

func (i *iter) Close() {
    i.key = nil
    i.value = nil
    i.current = nil
    i.list = nil
}
