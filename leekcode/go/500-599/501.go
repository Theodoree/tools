package _00_599


/*
501. 二叉搜索树中的众数

给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].
*/

var pre, maxCount, count int
var ret []int

func findMode(root *TreeNode) []int {
    if nil == root {
        return []int{}
    }
    pre, maxCount, count = 0, 0, 0
    transfer(root)
    return ret
}
func transfer(root *TreeNode) {
    if nil == root {
        return
    }
    transfer(root.Left)
    if pre == root.Val {
        count++
    } else {
        pre = root.Val
        count = 1
    }

    if count > maxCount {
        maxCount = count
        ret = []int{pre}
    } else if count == maxCount {
        ret = append(ret, pre)
    }
    transfer(root.Right)
}
