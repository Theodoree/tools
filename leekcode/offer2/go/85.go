package _go

/*
剑指 Offer II 085. 生成匹配的括号
正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/
func generateParenthesis(n int) []string {
	ret := []string{}

	backtrack(&ret, "", 0, 0, n)
	return ret

}

func backtrack(ret *[]string, ans string, left, right, n int) {
	if len(ans) == 2*n {
		*ret = append(*ret, ans)
		return
	}
	if left < n {
		backtrack(ret, ans+"(", left+1, right, n)
	}
	if right < left {
		backtrack(ret, ans+")", left, right+1, n)
	}
}
