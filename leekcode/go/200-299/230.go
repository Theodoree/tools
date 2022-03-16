package _00_299



/*
230. 二叉搜索树中第K小的元素

给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3
进阶：
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？

func kthSmallest(root *TreeNode, k int) int {
    var st []*TreeNode
    p := root
    count := 0
    for p != nil || len(st) >0 {
        if p != nil{
            st = append(st,p)
            p = p.Left
        } else{
            count++
            if(count == k){
                return st[len(st)-1].Val
            }
            p = st[len(st)-1].Right
            st = st[:len(st)-1]
        }
    }
    return 0;
}
*/
func kthSmallest(root *TreeNode, k int) int {
    val := []int{}
    GetSlice(root, &val, k)

    return val[k-1]
}

func GetSlice(root *TreeNode, val *[]int, k int) {

    if root == nil || len(*val) >= k {
        return
    }

    GetSlice(root.Left, val, k)
    *val = append(*val, root.Val)
    GetSlice(root.Right, val, k)

}