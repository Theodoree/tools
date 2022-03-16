package _600_1699



/*
1609. 奇偶树
如果一棵二叉树满足下述几个条件，则可以称为 奇偶树 ：
二叉树根节点所在层下标为 0 ，根的子节点所在层下标为 1 ，根的孙节点所在层下标为 2 ，依此类推。
偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减
给你二叉树的根节点，如果二叉树为 奇偶树 ，则返回 true ，否则返回 false 。
*/
func isEvenOddTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	for level := 0; len(q) > 0; level ^= 1 {
		prev := 0
		if level == 1 {
			prev = math.MaxInt32
		}
		tmp := q
		q = nil
		for _, node := range tmp {
			val := node.Val
			if val % 2 == level || val % 2 == 0 && prev <= val || val % 2 == 1 && prev >= val {
				return false
			}
			prev = val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return true
}

func isEvenOddTree(root *TreeNode) bool {
	var result [][]int
	levelOrder(root,0,&result)
	for i:=0;i<len(result);i++{
		flag :=i%2==0
		for j:=0;j<len(result[i]);j++{
			if flag {
				if j+1 <  len(result[i]) && result[i][j] >= result[i][j+1] {
					return false
				}
				if result[i][j]%2 == 0 {
					return false
				}
			}

			if !flag {
				if j+1 <  len(result[i]) && result[i][j] <= result[i][j+1] {
					return false
				}
				if result[i][j]%2 == 1 {
					return false
				}
			}
		}
	}
	return true

}
func levelOrder(root *TreeNode,depth int,result *[][]int) {
	if root == nil {
		return
	}
	if depth >= len(*result){
		*result =append(*result,[]int{})
	}
	(*result)[depth] = append((*result)[depth],root.Val)
	levelOrder(root.Left,depth+1,result)
	levelOrder(root.Right,depth+1,result)
}

