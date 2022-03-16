package _go

/*
面试题 08.09. 括号
括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。

说明：解集不能包含重复的子集。

例如，给出 n = 3，生成结果为：
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
