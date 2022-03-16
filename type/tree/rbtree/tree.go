package rbtree

import "fmt"

/*
https://zh.wikipedia.org/wiki/%E7%BA%A2%E9%BB%91%E6%A0%91
*/

/*
   rule-1:节点是红色或黑色。
   rule-2:根是黑色。
   rule-3:所有叶子都是黑色（叶子是NIL节点）。
   rule-4:每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
   rule-5:从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
*/

const (
    leftNodeIndex  = 0
    rightNodeIndex = 1
)

type RBTree struct {
    root *node
    pool *resourcesPool
    size uint64
}

func NewRBTree() *RBTree {
    return &RBTree{
        root: nil,
        pool: _resourcesPool,
        size: 0,
    }
}

/*
desc:返回size
*/
func (tree *RBTree) Size() uint64 {
    return tree.size
}

/*
desc:插入k、v
*/
func (tree *RBTree) Insert(key, value uint64) {
    /*
       case:如果root为空,先插入root(性质2)
    */
    if tree.root == nil {
        tree.size++
        tree.newRoot(key, value)
        return
    }
    /*
       case:插入root
    */
    n := tree.newNode(key, value)
    if tree.insert(tree.root, n) {
        tree.size++
    }
}

/*
desc:插入k、v
*/
func (tree *RBTree) insert(parent, n *node) bool {
    /*
       case:左孩子必然小于parent,右孩子必然大于parent
    */
    if parent.key > n.key {
        if parent.child[leftNodeIndex] != nil {
            tree.insert(parent.child[leftNodeIndex], n)
        } else {
            parent.child[leftNodeIndex] = n
            n.parent = parent
            tree.fixTree(n)
        }
    } else if parent.key < n.key {
        if parent.child[rightNodeIndex] != nil {
            tree.insert(parent.child[rightNodeIndex], n)
        } else {
            parent.child[rightNodeIndex] = n
            n.parent = parent
            tree.fixTree(n)
        }
    } else { /* 已经存在 */
        parent.value = n.value
        tree.pool.putNode(n)
        return false
    }
    return true
}

/*
desc:维护红黑树性质
*/
func (tree *RBTree) fixTree(n *node) {

    /*
       case-1: 新节点n位于树的根上,没有父节点,设置为黑色(规则2)
    */
case1:
    if n.parent == nil {
        n.color = black
        tree.root = n
        return
    }

    /*
       case-2: 新节点必定红色的,子节点是黑色,如果父节点也是黑色的,那么就满足(规则-4)
    */
    if n.parent.color == black {
        return
    }

    /*
       case-3:
    */

    if n.Uncle() != nil && n.Uncle().color == red {
        n.parent.color = black
        n.Uncle().color = black
        n.GrandParent().color = red
        goto case1
    }

    /*
       case-4:
    */

    if n == n.parent.child[rightNodeIndex] && n.parent == n.GrandParent().child[leftNodeIndex] {
        tree.rotateLeft(n)
        n = n.child[leftNodeIndex]
    } else if n == n.parent.child[leftNodeIndex] && n.parent == n.GrandParent().child[rightNodeIndex] {
        tree.rotateRight(n)
        n = n.child[rightNodeIndex]
    }

    /*
       case-5 :
    */
    n.parent.color = black
    n.GrandParent().color = red

    if n == n.parent.child[leftNodeIndex] && n.parent == n.GrandParent().child[leftNodeIndex] {
        tree.rotateRight(n.parent)
    } else {
        tree.rotateLeft(n.parent)
    }

    /*
       rule-1:节点是红色或黑色。
       rule-2:根是黑色。
       rule-3:所有叶子都是黑色（叶子是NIL节点）。
       rule-4:每个红色节点必须有两个黑色的子节点。（从每个叶子到根的所有路径上不能有两个连续的红色节点。）
       rule-5:从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点。
    */

}

/*
desc:插入新root
*/
func (tree *RBTree) newRoot(key, value uint64) {
    tree.root = tree.newNode(key, value)
    tree.root.color = black
}

func (tree *RBTree) newNode(key, value uint64) *node {
    node := tree.pool.getNode()
    node.key = key
    node.value = value
    return node
}

/*
desc:删除给定key的节点
*/
func (tree *RBTree) Delete(key uint64) bool {
    if tree.delete(tree.root, key) {
        tree.size--
        return true
    }
    return false

}

/*
desc:删除给定key的节点
*/

func (tree *RBTree) delete(n *node, key uint64) bool {
    /*
       case:n.value > key,那么选择更小的节点(左孩子比父节点小)
       case:n.value < key,那么选择更大的节点(右孩子比父节点大)
       case:等于,那么就找到了要删除的节点
    */

    if n.value > key {
        if n.child[leftNodeIndex] == nil {
            return false
        }
        return tree.delete(n.child[leftNodeIndex], key)
    } else if n.value < key {
        if n.child[rightNodeIndex] == nil {
            return false
        }
        return tree.delete(n.child[rightNodeIndex], key)
    } else {
        /*
           case:如果右节点为空,那么n节点就只有一个或者零个子节点
        */
        if n.child[rightNodeIndex] == nil {
            tree.deleteChild(n)
            return true
        }

        /*
           case:获取右子树中最小的值替换父节点
        */

        small := tree.getMinNode(n.child[rightNodeIndex])
        n.value = small.value
        n.key = small.key
        tree.deleteChild(small)
        return true
    }

}

/*
desc:删除
*/
func (tree *RBTree) deleteChild(p *node) {
    /*
       case:删除的节点是根节点且没有子节点
    */
    if p.parent == nil && p.child[leftNodeIndex] == nil && p.child[rightNodeIndex] == nil {
        p = nil
        tree.root = nil
        return
    }


    var child = p.child[leftNodeIndex]
    if child == nil {
        child = p.child[rightNodeIndex]
    }

    /*
       case:如果删除的节点是根节点,那么将子节点提拔为root且要满足规则-2
    */
    if p.parent == nil {
        child.parent = nil
        tree.root = child
        tree.root.color = black
        return
    }

    /*
       case:关联祖先节点和child节点
    */

    if p.parent.child[leftNodeIndex] == p {
        p.parent.child[leftNodeIndex] = child
    } else {
        p.parent.child[rightNodeIndex] = child
    }

    child.parent = p.parent

    /*
       case:
    */
    if p.color == black {
        if child.color == red {
            child.color = black
        } else {
            tree.deleteFixTree(child)
        }
    }

}

func (tree *RBTree) deleteFixTree(p *node) {

    /*
       case-1:
    */
case1:
    if p.parent == nil {
        p.color = black
        return
    }
    /*
       case-2:
    */
    s := p.sibling()
    if s.color == red {
        p.parent.color = red
        p.color = black
        if p == p.parent.child[leftNodeIndex] {
            tree.rotateLeft(p.parent)
        } else {
            tree.rotateRight(p.parent)
        }
    }

    /*
       case-3:
    */
    s = p.sibling()
    if p.parent.color == black && s.color == black && s.child[leftNodeIndex].color == black && s.child[rightNodeIndex].color == black {
        s.color = red
        p = p.parent
        goto case1
    }

    /*
       case-4:
    */

    s = p.sibling()

    if p.parent.color == red &&
        s.color == black &&
        s.child[leftNodeIndex].color == black &&
        s.child[rightNodeIndex].color == black {
        s.color = red
        p.parent.color = black
        return
    }

    /*
       case-5:
    */
    if s.color == black {
        if p == p.parent.child[leftNodeIndex] && s.child[rightNodeIndex].color == black && s.child[leftNodeIndex].color == red {
            s.color = red
            s.child[leftNodeIndex].color = black
            tree.rotateRight(s)
        } else if p == p.parent.child[rightNodeIndex] && s.child[leftNodeIndex].color == black && s.child[rightNodeIndex].color == red {
            s.color = red
            s.child[rightNodeIndex].color = black
            tree.rotateLeft(s)
        }
    }

    /*
       case-6:
    */

    s = p.sibling()

    s.color = p.parent.color
    p.parent.color = black

    if p == p.parent.child[leftNodeIndex] {
        p.child[rightNodeIndex].color = black
        tree.rotateLeft(p.parent)
    } else {
        s.child[leftNodeIndex].color = black
        tree.rotateRight(p.parent)
    }

}

/*
desc:返回给定Key的value
*/
func (tree *RBTree) Find(key uint64) (uint64, bool) {
    n := tree.root
    for n != nil {
        /*
           case: n.key == key,那么直接返回
           case: m.key > key,那么选择一个左孩子继续比较(左孩子必然比父节点小)
           case: m.key < key,那么选择一个右孩子(右孩子必定比父节点大)
        */
        switch {
        case n.key == key:
            return n.value, true
        case n.key > key:
            n = n.child[leftNodeIndex]
        case n.key < key:
            n = n.child[rightNodeIndex]
        }

    }

    return 0, false
}

/*
desc:左旋
*/
func (tree *RBTree) rotateLeft(p *node) {
    if p.parent == nil {
        tree.root = p
    }
    gp := p.GrandParent()
    fa := p.parent
    left := p.child[leftNodeIndex]

    fa.child[rightNodeIndex] = left

    if left != nil {
        left.parent = fa
    }
    p.child[rightNodeIndex] = fa
    fa.parent = p

    if tree.root == fa {
        tree.root = p
    }

    if gp != nil {
        if gp.child[leftNodeIndex] == fa {
            gp.child[leftNodeIndex] = p
        } else {
            gp.child[rightNodeIndex] = p
        }
    }

}

/*
desc:右旋
*/
func (tree *RBTree) rotateRight(p *node) {

    gp := p.GrandParent()
    fa := p.parent
    right := p.child[rightNodeIndex]

    fa.child[leftNodeIndex] = right

    if right != nil {
        right.parent = fa
    }
    p.child[rightNodeIndex] = fa
    fa.parent = p

    if tree.root == fa {
        tree.root = p
    }
    p.parent = gp

    if gp != nil {
        if gp.child[leftNodeIndex] == fa {
            gp.child[leftNodeIndex] = p
        } else {
            gp.child[rightNodeIndex] = p
        }
    }

}

func (tree *RBTree) getMinNode(n *node) *node {
    for n.child[leftNodeIndex] != nil {
        n = n.child[leftNodeIndex]
    }
    return n
}

func (tree *RBTree) LDR() {
    ldr(tree.root, 0)
}

func ldr(n *node, depth uint) {
    if n == nil {
        return
    }
    ldr(n.child[leftNodeIndex], depth+1)
    fmt.Printf("key %d value %d depth %d", n.key, n.value, depth)
    ldr(n.child[rightNodeIndex], depth+1)
}
