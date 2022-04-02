package _100_2199


/*
2196. 根据描述创建二叉树
给你一个二维整数数组 descriptions ，其中 descriptions[i] = [parenti, childi, isLefti] 表示 parenti 是 childi 在 二叉树 中的 父节点，二叉树中各节点的值 互不相同 。此外：
如果 isLefti == 1 ，那么 childi 就是 parenti 的左子节点。
如果 isLefti == 0 ，那么 childi 就是 parenti 的右子节点。
请你根据 descriptions 的描述来构造二叉树并返回其 根节点 。
测试用例会保证可以构造出 有效 的二叉树。
*/

func createBinaryTree(descriptions [][]int) *TreeNode {

	var m = make(map[int]bool)

	var treeMap = make(map[int]*TreeNode)
	for _, desc := range descriptions {
		parent := desc[0]
		p := treeMap[parent]
		if _, ok := m[parent]; !ok {
			m[parent] = true
		}
		if p == nil {
			p = &TreeNode{
				Val: parent,
			}
			treeMap[parent] = p
		}

		node := treeMap[desc[1]]
		if node == nil {
			node = &TreeNode{
				Val: desc[1],
			}
			treeMap[desc[1]] = node
		}

		if desc[2] == 1 { // left
			p.Left = node
		} else { // right
			p.Right = node
		}
		m[desc[1]] = false
	}

	for val, ok := range m {
		if ok {
			return treeMap[val]
		}
	}

	return nil

}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := make([]*TreeNode, 1e5 + 1)
	hasParent := make([]bool, 1e5 + 1)
	for _, description := range descriptions {
		parent, child, isLeft := description[0], description[1], description[2]
		if nodes[child] == nil {
			nodes[child] = &TreeNode{
				Val: child,
			}
		}
		if nodes[parent] == nil {
			nodes[parent] = &TreeNode{
				Val: parent,
			}
		}
		hasParent[child] = true
		if isLeft == 1 {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}
	}
	for root := 1; root <= 1e5; root++ {
		if nodes[root] != nil && !hasParent[root] {
			return nodes[root]
		}
	}
	return nil
}

