package _00_999


/*
988. 从叶结点开始的最小字符串
给定一颗根结点为 root 的二叉树，树中的每一个结点都有一个从 0 到 25 的值，分别代表字母 'a' 到 'z'：值 0 代表 'a'，值 1 代表 'b'，依此类推。
找出按字典序最小的字符串，该字符串从这棵树的一个叶结点开始，到根结点结束。
（小贴士：字符串中任何较短的前缀在字典序上都是较小的：例如，在字典序上 "ab" 比 "aba" 要小。叶结点是指没有子结点的结点。）
*/

func smallestFromLeaf(root *TreeNode) string {
	var min string
	dfs(root,"",&min)
	return min
}

func dfs(root *TreeNode,str string,min *string) {
	if root == nil {
		return
	}

	str +=string(byte(root.Val)+'a')
	if root .Left == nil && root.Right == nil {
		var f =make([]byte,0,len(str))
		for i:=len(str)-1;i>=0;i--{
			f = append(f,str[i])
		}
		str = string(f)
		if *min == "" {
			*min = str
		}else if str < *min {
			*min = str
		}

	}

	dfs(root.Left,str,min)
	dfs(root.Right,str,min)

}

