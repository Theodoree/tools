package _00_699


/*
653. 两数之和 IV - 输入 BST

给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

案例 1:

输入:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

输出: True


案例 2:

输入:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

输出: False
提交记录
*/

func findTarget(root *TreeNode, k int) bool {

    var result []int

    _LDR(root, &result)

    var left, right int
    right = len(result) - 1
    for left < right {
        if result[left]+result[right] > k {
            right--
        } else if result[left]+result[right] < k {
            left++
        } else {
            return true
        }
    }

    return false
}

func _LDR(root *TreeNode, result *[]int) {
    if root == nil {
        return
    }

    _LDR(root.Left, result)
    *result = append(*result, root.Val)
    _LDR(root.Right, result)

}

