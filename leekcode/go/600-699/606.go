package _00_699

import (
    "strconv"
    "strings"
)

/*
606. 根据二叉树创建字符串

你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。

空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

示例 1:

输入: 二叉树: [1,2,3,4]
       1
     /   \
    2     3
   /
  4

输出: "1(2(4))(3)"

解释: 原本将是“1(2(4)())(3())”，
在你省略所有不必要的空括号对之后，
它将是“1(2(4))(3)”。
示例 2:

输入: 二叉树: [1,2,3,null,4]
       1
     /   \
    2     3
     \
      4

输出: "1(2()(4))(3)"

解释: 和第一个示例相似，
除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。
*/

func tree2str(t *TreeNode) string {
    var b strings.Builder
    tree2strHelper(&b, t)
    return b.String()
}

func tree2strHelper(b *strings.Builder, t *TreeNode) {
    if t == nil {
        return
    }
    b.WriteString(strconv.Itoa(t.Val))

    if t.Left == nil && t.Right == nil {
        return
    }

    b.WriteString("(")
    tree2strHelper(b, t.Left)
    b.WriteString(")")

    if t.Right != nil {
        b.WriteString("(")
        tree2strHelper(b, t.Right)
        b.WriteString(")")
    }
}

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}