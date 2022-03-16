package skiplist

type Iterator interface {
    Next() bool                // 后继
    Previous() bool            // 前驱
    Key() interface{}          // 当前的key
    Value() interface{}        // 当前的value
    Seek(key interface{}) bool // 光标位置
    Close()                    // 关闭相关资源
}




