package _00_199


/*
133. 克隆图
给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。

图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。
class Node {
    public int val;
    public List<Node> neighbors;
}
*/

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	var n Node
	arr:=make([]*Node,101)
	dfs(node,arr,&n)
	return &n
}

func dfs(node *Node,arr []*Node,clone *Node){
	clone.Val = node.Val
	arr[clone.Val] = clone
	for _,v:=range node.Neighbors{
		var t *Node
		if arr[v.Val] != nil {
			t =  arr[v.Val]
		}else{
			t = &Node{Val: v.Val}
			arr[v.Val] = t
			dfs(v, arr, t)
		}
		clone.Neighbors = append(clone.Neighbors,t)
	}
}

