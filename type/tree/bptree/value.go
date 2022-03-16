package bptree

type value struct {
    val uint64
}

func NewValue() *value {
    return &value{}
}

func (v *value) reset() {
    v.val = 0
}
