package _00_199


/*
173. 二叉搜索树迭代器

实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。

调用 next() 将返回二叉搜索树中的下一个最小的数。



示例：



BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false

提示：

next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
*/

type BSTIterator struct {
    arr []int
}

func Constructor(root *TreeNode) BSTIterator {
    if root == nil {
        return BSTIterator{}
    }
    var arr []int
    DFS(root, &arr) // 仔细阅读了题目,发现不一定要用最小堆做,因为没有插入,而且这棵树是二叉搜索树,所以直接中序遍历

    return BSTIterator{arr: arr}
}

func DFS(root *TreeNode, cur *[]int) {
    if root == nil {
        return
    }

    DFS(root.Left, cur)
    * cur = append(*cur, root.Val)
    DFS(root.Right, cur)

}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    if !this.HasNext() {
        return -1
    }
    val := this.arr[0]
    this.arr = this.arr[1:]
    return val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.arr) > 0
}

