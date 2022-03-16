package _000_1099

import "sort"

/*
1079. 活字印刷
你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。
注意：本题中，每个活字字模只能使用一次。
*/
func numTilePossibilities(tiles string) int {
	var m = make(map[string]struct{},100)
	Backtracking(tiles,0,"",m) // 典中典之指数阶
	return len(m)
}
func Backtracking(t string,flag int,cur string,m map[string]struct{}){
	if flag > 0 {
		m[cur] = struct{}{}
	}

	for i:=0;i<len(t);i++{
		if (1<<i) & flag  > 0 {
			continue
		}

		Backtracking(t,flag|1<<i,cur+string(t[i]),m)
	}
}


func numTilePossibilities(tiles string) int {
	chs := []byte(tiles)
	sort.Slice(chs, func(i, j int) bool {
		return chs[i] < chs[j]
	})
	n := len(chs)
	res := 0
	used := make([]bool, n)
	var dfs func()
	dfs = func() {
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && chs[i] == chs[i - 1] && !used[i - 1] {
				continue
			}
			used[i] = true
			res++
			dfs()
			used[i] = false
		}
	}
	dfs()
	return res
}