package _00_399

func verticalOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	virtualValNodesRel := map[int][]int{}
	q := []*TreeNodeVal{{root, 0}}
	minKey, maxKey := 0, 0
	for len(q) > 0 {

		elem := q[0]
		q = q[1:]

		virtualVal := elem.val
		minKey = min(minKey, virtualVal)
		maxKey = max(maxKey, virtualVal)
		virtualValNodesRel[virtualVal] = append(virtualValNodesRel[virtualVal], elem.node.Val)

		if elem.node.Left != nil {
			q = append(q, &TreeNodeVal{elem.node.Left, virtualVal - 1})
		}
		if elem.node.Right != nil {
			q = append(q, &TreeNodeVal{elem.node.Right, virtualVal + 1})
		}
	}

	for i := minKey; i <= maxKey; i++ {
		ans = append(ans, virtualValNodesRel[i])
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
