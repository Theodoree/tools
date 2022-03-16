package _00_199

func maxDepth(root *TreeNode) int {
    var maxInt int
    MaxDepth(root, &maxInt, 0)
    return maxInt
}

func MaxDepth(node *TreeNode, maxInt *int, current int)  {
    if node == nil {
        if *maxInt < current {
            *maxInt = current
        }
        return
    }

     MaxDepth(node.Left, maxInt, current+1)
     MaxDepth(node.Right, maxInt, current+1)

}

/* 另外一个答案
func maxDepth(root *TreeNode) int {
    if nil == root {
        return 0
    }
    maxLeft := maxDepth(root.Left)
    maxRight := maxDepth(root.Right)
    if maxLeft >= maxRight {
        return maxLeft + 1
    } else {
        return maxRight + 1
    }
}

*/