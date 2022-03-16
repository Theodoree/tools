package _00_899


/*
897. 递增顺序查找树

给定一个树，按中序遍历重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。



示例 ：

输入：[5,3,6,2,4,null,8,1,null,null,null,7,9]

       5
      / \
    3    6
   / \    \
  2   4    8
 /        / \
1        7   9

输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

 1
  \
   2
    \
     3
      \
       4
        \
         5
          \
           6
            \
             7
              \
               8
                \
                 9
*/

func increasingBST(root *TreeNode) *TreeNode {

    if root == nil {
        return nil
    }

    head := new(TreeNode)
    result := head
    IncreasingBST(root, result)

    return head.Right
}

func IncreasingBST(root *TreeNode, result *TreeNode) *TreeNode {
    if root == nil {
        return result
    }
    result = IncreasingBST(root.Left, result)
    result.Right = &TreeNode{Val: root.Val}
    result = result.Right
    result = IncreasingBST(root.Right, result)
    return result

}
