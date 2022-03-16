package _go

import (
	"strconv"
	"strings"
)

/*
剑指 Offer II 048. 序列化与反序列化二叉树
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
*/
type Codec struct {
}

func Constructor() Codec {
	return Codec{}

}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var buf []string
	maxDepth := getMaxDepth(root)
	LevelTree(root, &buf, 0, maxDepth)

	return strings.Join(buf, ",")
}

func LevelTree(root *TreeNode, result *[]string, depth int, maxDepth int) {

	if depth < maxDepth && len(*result) <= depth {
		*result = append(*result, "")
	}
	k := *result

	if root == nil {
		if depth < maxDepth {
			if len(k[depth]) > 0 {
				k[depth] += "," + Null
			} else {
				k[depth] = Null
			}

		}
		return
	}

	if len(k[depth]) > 0 {
		k[depth] += "," + strconv.Itoa(root.Val)
	} else {
		k[depth] = strconv.Itoa(root.Val)
	}
	LevelTree(root.Left, result, depth+1, maxDepth)
	LevelTree(root.Right, result, depth+1, maxDepth)

}

func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	a := getMaxDepth(root.Left) + 1
	b := getMaxDepth(root.Right) + 1
	if a > b {
		return a
	}
	return b
}

var Null = "null"

func CreateTree(b []string) *TreeNode {

	var root *TreeNode
	if len(b) == 0 {
		return nil
	}
	root = &TreeNode{Val: atoi(b[0])}

	v := make([]*TreeNode, len(b)+1)
	var index, index1 int
	index = 1
	index1 = 1
	current := root
	v[0] = root
	for i := 1; i <= len(b)/2; i++ {

		if b[i*2-1] != Null {
			current.Left = &TreeNode{Val: atoi(b[i*2-1])}
			v[index1] = current.Left
			index1++
		}
		if i*2 < len(b) && b[i*2] != Null {
			current.Right = &TreeNode{Val: atoi(b[i*2])}
			v[index1] = current.Right
			index1++
		}
		current = v[index]
		index++
	}
	return root

}

func atoi(val string) int {
	num, _ := strconv.Atoi(val)
	return num
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	return CreateTree(strings.Split(data, ","))
}
