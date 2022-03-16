package _000_1099


/*
1023. 驼峰式匹配
如果我们可以将小写字母插入模式串 pattern 得到待查询项 query，那么待查询项与给定模式串匹配。（我们可以在任何位置插入每个字符，也可以插入 0 个字符。）
给定待查询列表 queries，和模式串 pattern，返回由布尔值组成的答案列表 answer。只有在待查项 queries[i] 与模式串 pattern 匹配时， answer[i] 才为 true，否则为 false。
*/

func camelMatch(queries []string, p string) []bool {
	ans := []bool{}
	for _, q := range queries {
		ans = append(ans, isMatch(q, p))
	}
	return ans
}

func isMatch(q, p string) bool {
	//最后添加一个特殊字符串，因q中不可能出现特殊字符串所P-1匹配完后不可能匹配到
	p = p + "$"
	pi := 0
	for qi := range q {
		if q[qi] == p[pi] {
			pi++
		} else if q[qi] <= 'Z' {
			return false
		}
	}
	return pi == len(p)-1
}




