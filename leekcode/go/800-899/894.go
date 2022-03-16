package _00_899

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func allPossibleFBT(N int) []*TreeNode {

    result := []*TreeNode{}

    AllPossibleFBT(nil, nil,8, 0, &result)
    fmt.Println(result)
    return result
}

func AllPossibleFBT(current *TreeNode,Head *TreeNode, N, c int, result *[]*TreeNode) {
    node := &TreeNode{Val: 0}

    if c == N{
        *result = append(*result,Head)
    }


    if current == nil {
        AllPossibleFBT(node,node, N, 1, result)
    }else{
        current.Left = node
        AllPossibleFBT(node,Head, N, c+1, result)
        current.Right = node
        AllPossibleFBT(node,Head, N, c+1, result)
    }
}

func CheckNode(node *TreeNode) bool {

    if node.Left == nil && node.Right == nil {
        return true
    }

    if node.Left != nil && node.Right != nil {
        CheckNode(node.Left)
        CheckNode(node.Right)
    } else {
        return false
    }
    return true
}

func main() {
    allPossibleFBT(8)
}
