package avltree

/*
https://zh.wikipedia.org/wiki/AVL%E6%A0%91
*/

const (
    leftNodeIndex  = 0
    rightNodeIndex = 1
    maxIndex       = 1
)

type Tree struct {
    root       *node      // 根
    comparator Comparator // 比较器
    size       int        // 树的节点数量
    pool       *resourcesPool
}

func newTree() *Tree {
    return &Tree{pool: _resourcesPool}
}

func NewWithIntComparator() *Tree {
    tree := newTree()
    tree.comparator = intComparator
    return tree
}

/*
  rule-A:最小节点永远在最左边,也就是索引0。
  rule-B:最大节点永远在最右边,也就是索引1
  rule-C:左孩子必然大于右孩子,父节点必定大于左孩子且小于右孩子
  rule-D:左右子树高度之差的绝对值不超过（-1 / 0 / 1）
  rule-E:左右子树必定也是AVT树
*/

func (tree *Tree) newNode(key, value interface{}, parent *node) *node {
    node := tree.pool.getNode()
    node.key = key
    node.value = value
    node.parent = parent
    return node
}

/*
desc:按顺序返回key值
*/
func (tree *Tree) Keys() []interface{} {
    keys := make([]interface{}, 0,tree.size)
    LDR(tree.root,&keys,nil)
    return keys
}
/*
desc:中序遍历
*/
func LDR(n *node, keyArr, valueArr *[]interface{}) {
    if n == nil {
        return
    }

    LDR(n.child[0],keyArr,valueArr)
    if keyArr != nil {
        *keyArr = append(*keyArr,n.key)
    }
    if valueArr != nil {
        *valueArr = append(*valueArr,n.value)
    }
    LDR(n.child[1],keyArr,valueArr)


}

/*
desc:按顺序返回value值
*/
func (tree *Tree) Values() []interface{} {
    values := make([]interface{}, 0,tree.size)
    LDR(tree.root,nil,&values)
    return values
}

/*
desc:找寻一个node
*/
func (tree *Tree) Get(key interface{}) (interface{}, bool) {
    n := tree.root

    for n != nil {
        cmp := tree.comparator(key, n.key)
        /*
           case: 如果符合条件,直接返回
           case: 如果key小于n.Key,那么选择左边节点(rule A),选择一个更小的节点比较
           case: 如果key大于n.Key,那么选择右边节点(rule B),选择一个更大的节点比较
        */
        switch {
        case cmp == 0:
            return n.value, true
        case cmp < 0:
            n = n.child[leftNodeIndex]
        case cmp > 0:
            n = n.child[rightNodeIndex]
        }
    }
    return nil, false
}

/*
desc:将root指向设置为空,那么gc就可以回收了
*/
func (tree *Tree) Clear() {
    tree.root = nil
    tree.size = 0
}
func (tree *Tree) Size() int {
    return tree.size
}
func (tree *Tree) Empty() bool {
    return tree.size == 0 || tree.root == nil
}
/*
desc:删除给定key的节点
*/
func (tree *Tree) Delete(key interface{}) {
    tree.delete(key, &tree.root)
}

/*
desc:获取树中最小节点(可能会为空)
*/
func (tree *Tree) Left() *node {
    return tree.bottom(leftNodeIndex)
}
/*
desc:获取树中最大节点(可能会为空)
*/
func (tree *Tree) Right() *node {
    return tree.bottom(rightNodeIndex)
}
/*
desc:根据入参的索引,循环寻找符合条件的node
*/
func (tree *Tree) bottom(idx uint64) *node {
    if idx > maxIndex || tree.Empty() {
        return nil
    }

    n := tree.root
    /*
       case:往底层树循环检查
    */
    for n.child[idx] != nil {
        n = n.child[idx]
    }
    return n
}


/*
desc:获取给定Key的node,如果没有则返回小于该key的父节点
*/
func (tree *Tree) Floor(key interface{}) (floor *node, found bool) {
    found = false
    n := tree.root
    for n != nil {
        c := tree.comparator(key, n.key)
        /*
           case: 如果符合条件,直接返回(必然有floor节点)
           case: 如果key小于n.Key,那么选择左边节点(rule A),选择一个更小的节点比较
           case: 如果key大于n.Key,那么选择右边节点(rule B),选择一个更大的节点比较(floor节点)
        */
        switch {
        case c == 0:
            return n, true
        case c < 0:
            n = n.child[0]
        case c > 0:
            floor, found = n, true
            n = n.child[1]
        }
    }
    if found {
        return
    }
    return nil, false
}

/*
desc:获取给定Key的node,如果没有则返回大于该key的ceil节点
*/
func (tree *Tree) Ceiling(key interface{}) (floor *node, found bool) {
    found = false
    n := tree.root
    for n != nil {
        c := tree.comparator(key, n.key)
        /*
           case: 如果符合条件,直接返回(必然有ceiling节点)
           case: 如果key小于n.Key,那么选择左边节点(rule A),选择一个更小的节点比较(ceiling节点)
           case: 如果key大于n.Key,那么选择右边节点(rule B),选择一个更大的节点比较
        */
        switch {
        case c == 0:
            return n, true
        case c < 0:
            floor, found = n, true
            n = n.child[0]
        case c > 0:
            n = n.child[1]
        }
    }
    if found {
        return
    }
    return nil, false
}

/*
desc:插入一个k、v
*/
func (tree *Tree) Put(key, value interface{}) {
    tree.put(key, value, nil, &tree.root)
}

/*
desc:插入一个k、v
*/
func (tree *Tree) put(key, value interface{}, parent *node, np **node) bool {
    n := *np

    /*
       case: 插入成功
    */
    if n == nil {
        tree.size++
        *np = tree.newNode(key, value, parent)
        return true
    }

    flag := tree.comparator(key, n.key)
    /*
       case: 如果找到了符合条件的节点直接更新
    */
    if flag == 0 {
        n.key = key
        n.value = value
        return false
    }

    if flag < 0 {
        flag = -1
    } else {
        flag = 1
    }
    /*
       0 = ( -1 + 1 ) / 2
       1 = ( 1 + 1 ) / 2
    */
    idx := (flag + 1) / 2

    /*
       case:递归寻找位置,如果插入成功,那么insertd为true
    */
    inserted := tree.put(key, value, n, &n.child[idx])
    if inserted {
        return tree.putFix(int8(flag), np)
    }

    return false
}

/*
desc:删除一个node
*/
func (tree *Tree) delete(key interface{}, np **node) bool {
    n := *np
    if n == nil {
        return false
    }

    flag := tree.comparator(key, n.key)
    /*
       case: 找了指定节点,进行删除
    */
    if flag == 0 {
        tree.size--
        /*
           case:如果右孩子为空,直接提拔左孩子为新的n节点,因为只有一个节点不需要调整(不会影响性质)
        */
        if n.child[1] == nil {
            if n.child[0] != nil {
                n.child[0].parent = n.parent
            }
            *np = n.child[0]

            n.reset()
            tree.pool.putNode(n)
            return true
        }

        /*
           case: 左右孩子都有,则需要调整,保证avt的性质不变
        */
        fix := tree.removeMin(&n.child[1], n)
        if fix {
            return tree.removeFix(-1, np)
        }
        return false
    }

    if flag < 0 {
        flag = -1
    } else {
        flag = 1
    }
    /*
       0 = ( -1 + 1 ) / 2
       1 = ( 1 + 1 ) / 2
    */
    idx := (flag + 1) / 2

    /*
       case: 递归删除
    */

    fix := tree.delete(key, &n.child[idx])
    if fix {
        return tree.removeFix(int8(-flag), np)
    }
    return false

}

func (tree *Tree) removeMin(np **node, parent *node) bool {
    n := *np
    /*
       case: 如果n的左孩子为空,那么直接将n的右孩子直接替换n,那么就删除成功了
    */
    if n.child[0] == nil {
        parent.key = n.key
        parent.value = n.value
        if n.child[1] != nil {
            n.child[1].parent = parent
        }
        *np = n.child[1]
        return true
    }
    fix := tree.removeMin(&n.child[0], n)
    if fix {
        return tree.removeFix(1, np)
    }
    return false
}

func (tree *Tree) putFix(c int8, np **node) bool {
    /*
    case:插入成功,那么需要维护该树的性质不变
    */
    s := *np
    if s.b == 0 {
        s.b = c
        return true
    }

    if s.b == -c {
        s.b = 0
        return false
    }

    if s.child[(c+1)/2].b == c {
        s = tree.singlerot(c, s)
    } else {
        s = tree.doublerot(c, s)
    }
    *np = s
    return false
}

func (tree *Tree) removeFix(c int8, np **node) bool {
    n := *np
    if n.b == 0 {
        n.b = c
        return false
    }

    if n.b == -c {
        n.b = 0
        return true
    }

    a := (c + 1) / 2
    if n.child[a].b == 0 {
        n = tree.rotate(c, n)
        n.b = -c
        *np = n
        return false
    }

    if n.child[a].b == c {
        n = tree.singlerot(c, n)
    } else {
        n = tree.doublerot(c, n)

    }
    *np = n
    return true
}

func (tree *Tree) singlerot(c int8, n *node) *node {
    n.b = 0
    n = tree.rotate(c, n)
    n.b = 0
    return n
}

func (tree *Tree) doublerot(c int8, n *node) *node {
    idx := (c + 1) / 2
    r := n.child[idx]
    n.child[idx] = tree.rotate(-c, n.child[idx])
    p := tree.rotate(c, n)

    switch {
    default:
        n.b = 0
        r.b = 0
    case p.b == c:
        n.b = -c
        r.b = 0
    case p.b == -c:
        n.b = 0
        r.b = c
    }

    p.b = 0
    return p
}

func (tree *Tree) rotate(c int8, n *node) *node {
    idx := (c + 1) / 2
    r := n.child[idx]
    n.child[idx] = r.child[idx^1]
    if n.child[idx] != nil {
        n.child[idx].parent = n
    }
    r.child[idx^1] = n
    r.parent = n.parent
    n.parent = r
    return r
}
