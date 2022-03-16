package skiplist

type Set struct {
    skiplist SkipList
}


func (s *Set) Add(key interface{}) {
    s.skiplist.Set(key, nil)
}

func (s *Set) Remove(key interface{}) (ok bool) {
    _, ok = s.skiplist.Delete(key)
    return ok
}

func (s *Set) Len() int {
    return s.skiplist.Len()
}


func (s *Set) Contains(key interface{}) bool {
    _, ok := s.skiplist.Get(key)
    return ok
}


func (s *Set) Iterator() Iterator {
    return s.skiplist.Iterator()
}

func (s *Set) Range(from, to interface{}) Iterator {
    return s.skiplist.Range(from, to)
}

func (s *Set) SetMaxLevel(newMaxLevel int) {
    s.skiplist.Level = newMaxLevel
}

// GetMaxLevel returns MaxLevel fo the underlying skip list.
func (s *Set) GetMaxLevel() int {
    return s.skiplist.Level
}

func NewSet() *Set {
    comparator := func(left, right interface{}) bool {
        return left.(Ordered).LessThan(right.(Ordered))
    }
    return NewCustomSet(comparator)
}

func NewCustomSet(lessThan func(l, r interface{}) bool) *Set {
    return &Set{skiplist: SkipList{
        lessThan: lessThan,
        header: &node{
            forward: []*node{nil},
        },
        Level: maxLevel,
    }}
}


func NewIntSet() *Set {
    return NewCustomSet(func(l, r interface{}) bool {
        return l.(int) < r.(int)
    })
}

func NewStringSet() *Set {
    return NewCustomSet(func(l, r interface{}) bool {
        return l.(string) < r.(string)
    })
}

