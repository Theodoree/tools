package _00_699


type Part struct {
	Node *TreeNode
	Number int
}

func widthOfBinaryTree(root *TreeNode) int {
	maxWidth := 0
	if root == nil{
		return maxWidth
	}
	queue := []Part{Part{root, 1}}
	for len(queue) > 0{
		size := len(queue)
		maxWidth = Max(maxWidth, queue[size-1].Number - queue[0].Number + 1)
		for i:=0;i<size;i++{
			node, num := queue[0].Node, queue[0].Number
			queue = queue[1:]
			if node.Left != nil{
				queue = append(queue, Part{node.Left, 2*num})
			}
			if node.Right != nil{
				queue = append(queue, Part{node.Right, 2*num+1})
			}
		}
	}
	return maxWidth
}


func Max(a, b int)int{
	if a > b{
		return a
	}
	return b
}

