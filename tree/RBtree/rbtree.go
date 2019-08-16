package RBtree

const (
	BLACK = true
	RED   = false
)

type RBTree struct {
	Root  *RBNode
	Count uint64
}

/*
红黑树的特性:
（1）每个节点或者是黑色，或者是红色。
（2）根节点是黑色。
（3）每个叶子节点（NIL）是黑色。 [注意：这里叶子节点，是指为空(NIL或NULL)的叶子节点！]
（4）如果一个节点是红色的，则它的子节点必须是黑色的。
（5）从一个节点到该节点的子孙节点的所有路径上包含相同数目的黑节点。

注意：
(01) 特性(3)中的叶子节点，是只为空(NIL或null)的节点。
(02) 特性(5)，确保没有一条路径会比其他路径长出俩倍。因而，红黑树是相对是接近平衡的二叉树。
*/

func (r *RBTree) IsRbTree() bool {
	//（1）每个节点或者是黑色，或者是红色。

	//（2）根节点是黑色。
	if r.Root.Color != BLACK {
		return false
	}

	//（3）每个叶子节点（NIL）是黑色。 [注意：这里叶子节点，是指为空(NIL或NULL)的叶子节点！]

	//（4）如果一个节点是红色的，则它的子节点必须是黑色的。

	//（5）从一个节点到该节点的子孙节点的所有路径上包含相同数目的黑节点。

	return true
}

func NewRBNode(key, value int64) *RBNode {
	return &RBNode{
		Key:   key,
		Value: value,
		Color: RED,
	}
}

func (r *RBTree) Insert(root *RBNode, key, value int64) *RBNode {
	if r.Root == nil {
		r.Count ++
		r.Root = NewRBNode(key, value)
		return r.Root
	}

	if root == nil {
		r.Count++
		node := NewRBNode(key, value)
		return node
	}
	if root.Key == key {
		root.Value = value
	} else if root.Key < key {
		root.Right = r.Insert(root.Right, key, value)
	} else if root.Key > key {
		root.Left = r.Insert(root.Left, key, value)
	}
	return root
}

/*
 * 对红黑树的节点(x)进行左旋转
 *
 * 左旋示意图(对节点x进行左旋)：
 *      px                              px
 *     /                               /
 *    x                               y
 *   /  \      --(左旋)-.           / \                #
 *  lx   y                          x  ry
 *     /   \                       /  \
 *    ly   ry                     lx  ly
 *
 *
 */

func (r *RBTree) LeftRotate(x *RBNode) {
	y := x.Right

	x.Right = x.Left
	if y.Left != nil {
		y.Left.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == nil {
		r.Root = y
	} else {
		if x.Parent.Left == x {
			x.Parent.Left = y
		} else {
			x.Parent.Right = y
		}
	}

	y.Left = x
	x.Parent = y

}

/*
 * 对红黑树的节点(y)进行右旋转
 *
 * 右旋示意图(对节点y进行左旋)：
 *            py                               py
 *           /                                /
 *          y                                x
 *         /  \      --(右旋)-.            /  \                     #
 *        x   ry                           lx   y
 *       / \                                   / \                   #
 *      lx  rx                                rx  ry
 *
 */
func (r *RBTree) RightRotate(y *RBNode) {
	x := y.Left
	y.Left = x.Right
	if x.Right != nil {
		x.Right.Parent = y
	}

	x.Parent = y.Parent

	if y.Parent == nil {
		r.Root = x
	} else {
		if y == y.Parent.Right {
			y.Parent.Right = x
		} else {
			y.Parent.Left = x
		}
	}

	x.Right = y
	y.Parent = x
}

func (r *RBTree) InsertFixUp(y *RBNode) {

	parent := y.GetParent()
	grandParent := y.GetGrandParent()

	// 若“父节点存在，并且父节点的颜色是红色”
	for parent != nil && parent.IsRed() {
		uncle, _ := y.GetUncle()
		//若“父节点”是“祖父节点的左孩子”
		if parent == grandParent.Left {
			// Case 1条件：叔叔节点是红色

			if uncle != nil && uncle.IsRed() {
				uncle.SetBlack()
				parent.SetBlack()
				grandParent.SetRed()
				y = grandParent
				continue
			}

			// Case 2条件：叔叔是黑色，且当前节点是右孩子
			if parent.Right == y {
				r.LeftRotate(parent)
				tmp := parent
				parent = y
				y = tmp
			}

			// Case 3条件：叔叔是黑色，且当前节点是左孩子。
			parent.SetBlack()
			grandParent.SetRed()
			r.RightRotate(grandParent)
		} else { //若“父节点”是“祖父节点的右孩子”

			// Case 1条件：叔叔节点是红色
			if uncle != nil && uncle.IsRed() {
				uncle.SetBlack()
			}
		}
	}

}

/*

            // Case 1条件：叔叔节点是红色
            RBTNode<T> uncle = gparent.left;
            if ((uncle!=null) && isRed(uncle)) {
                setBlack(uncle);
                setBlack(parent);
                setRed(gparent);
                node = gparent;
                continue;
            }

            // Case 2条件：叔叔是黑色，且当前节点是左孩子
            if (parent.left == node) {
                RBTNode<T> tmp;
                rightRotate(parent);
                tmp = parent;
                parent = node;
                node = tmp;
            }

            // Case 3条件：叔叔是黑色，且当前节点是右孩子。
            setBlack(parent);
            setRed(gparent);
            leftRotate(gparent);
        }
    }

    // 将根节点设为黑色
    setBlack(this.mRoot);
}
*/
