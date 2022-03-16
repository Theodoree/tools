package iterator

type User struct {
    Id uint64
}

// 其实像go map的迭代器,是有快照,但只是引用一个指针,当扩容发生时,那么桶内的元素是确定的,如果没有发生扩容,你还在里面进行添加的话的话,那么有意思的事情就发生了。元素会越遍历越多,
// 当然元素要插入到当前遍历的桶下标之后的桶。
type iterator interface {
    HasNext() bool
    Next() interface{}
}

type userIterator struct {
    arr      []User
    curIndex int
}

func (i *userIterator) HasNext() bool {
    return i.curIndex+1 <= len(i.arr)
}

func (i *userIterator) Next() (val interface{}) {
    if i.curIndex < len(i.arr) {
        val = i.arr[i.curIndex]
        i.curIndex++
    }
    return
}

func CreateUserIterator(arr []User) iterator {
    return &userIterator{
        arr:      arr,
        curIndex: 0,
    }
}
