package _200_1299

import "sort"

type trie struct {
	child [26]*trie
	end   bool
}

func (t *trie) Insert(s string)  {
	cur := t
	for _, v := range s {
		if cur.child[v-'a'] == nil {
			cur.child[v-'a'] = &trie{}
		}
		cur = cur.child[v-'a']
	}
	cur.end = true
}

func (t *trie) find(str string,cnt int) []string{
	node:=t
	for _,v:=range str{
		node = node.child[v-'a']
		if node == nil {
			return nil
		}
	}

	result:=make([]string,0,cnt)
	if node.end {
		result =append(result,str)
	}


	dfs(node,[]byte(str),&result,cnt)
	return result
}

func dfs(n *trie,baseString []byte,result *[]string,cnt int){
	if len(*result) >= cnt {
		return
	}
	for idx,v:=range n.child{
		if v != nil  && len(*result) < cnt {
			str:=make([]byte,len(baseString))
			copy(str,baseString)
			str = append(str,'a'+byte(idx))
			if v.end{
				*result = append(*result,string(str))
			}
			dfs(v,str,result,cnt)
		}
	}


}


/*
1268. 搜索推荐系统
给你一个产品数组 products 和一个字符串 searchWord ，products  数组中每个产品都是一个字符串。
请你设计一个推荐系统，在依次输入单词 searchWord 的每一个字母后，推荐 products 数组中前缀与 searchWord 相同的最多三个产品。如果前缀相同的可推荐产品超过三个，请按字典序返回最小的三个。
请你以二维列表的形式，返回在输入 searchWord 每个字母后相应的推荐产品的列表。
*/

func suggestedProducts(products []string, searchWord string) [][]string {
	var t trie
	for i:=0;i<len(products);i++{
		t.Insert(products[i])
	}

	var result [][]string
	for idx:=range searchWord{
		arr:=t.find(searchWord[:idx+1],3)
		result = append(result,arr)
	}
	return result
}

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	var ans [][]string
	for i:=0; i<len(searchWord); i++ {
		temp := searchWord[0:i+1]
		line := []string{}
		for j:=0; j<len(products); j++ {
			if len(products[j]) >= len(temp) && temp == products[j][0:i+1] {
				line = append(line, products[j])
			}
			if len(line) == 3 {
				break
			}
		}
		ans = append(ans, line)
	}

	return ans

}
