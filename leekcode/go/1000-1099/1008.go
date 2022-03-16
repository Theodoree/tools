package main

/*
1008. 先序遍历构造二叉树
返回与给定先序遍历 preorder 相匹配的二叉搜索树（binary search tree）的根结点。

(回想一下，二叉搜索树是二叉树的一种，其每个节点都满足以下规则，对于 node.left 的任何后代，
值总 < node.val，而 node.right 的任何后代，值总 > node.val。此外，先序遍历首先显示节点的值，
然后遍历 node.left，接着遍历 node.right。）

 

示例：

输入：[8,5,1,7,10,12]
输出：[8,5,10,1,7,null,12]

 

提示：
1 <= preorder.length <= 100

*/

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
    // DLR
    // LDR
    // DLR

    if len(preorder) == 0 {
        return nil
    }

    root := &TreeNode{Val: preorder[0]}

    current := root

    for i := 1; i < len(preorder); i++ {
        INSERT(current, preorder[i])
    }

    return root
}

func INSERT(root *TreeNode, val int) {
    node := &TreeNode{Val: val}
    for {
        if root.Val > val {
            if root.Left == nil {
                root.Left = node
                break
            }
            root = root.Left

        } else if root.Val < val {
            if root.Right == nil {
                root.Right = node
                break
            }
            root = root.Right
        } else {
            break
        }

    }

}

func DLR(root *TreeNode) {

    if root == nil {
        fmt.Println("null")
        return
    }

    fmt.Println(root.Val)
    DLR(root.Left)
    DLR(root.Right)
}


func main() {
    v := []int{8, 5, 1, 7, 10, 12}
    r := bstFromPreorder(v)

    DLR(r)

}