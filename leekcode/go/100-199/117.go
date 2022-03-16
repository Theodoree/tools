package _00_199

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var fn func(r *Node, result *[][]*Node, depth int)

	fn = func(r *Node, result *[][]*Node, depth int) {
		if r == nil {
			return
		}

		if len(*result) <= depth {
			*result = append(*result, []*Node{})
		}
		(*result)[depth] = append((*result)[depth], r)

		fn(r.Left, result, depth+1)
		fn(r.Right, result, depth+1)
	}

	root.Next = nil
	var result [][]*Node
	fn(root, &result, 0)

	for i := 1; i < len(result); i++ {
		for j := 0; j < len(result[i])-1; j++ {
			result[i][j].Next = result[i][j+1]
		}
	}

	return root
}
