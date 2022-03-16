package _00_499


/*
449. 序列化和反序列化二叉搜索树
序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。
编码的字符串应尽可能紧凑。
*/

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}
// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	buf,_:=json.Marshal(&root) // 可以没必要,可以考虑直接把树遍历出来,然后直接用一个数组装着
	return BytesToString(buf)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var n *TreeNode
	json.Unmarshal([]byte(data),n)
	return n
}


