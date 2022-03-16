package _go


/*
剑指 Offer II 086. 分割回文子字符串
给定一个字符串 s ，请将 s 分割成一些子串，使每个子串都是 回文串 ，返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。

*/

func partition(s string) [][]string {
	n := len(s)
	var res [][]string
	var tmp []string
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			res = append(res, append([]string{}, tmp...))
			return
		}
		for i := idx + 1; i <= n; i++ {
			if isPalindrome(s[idx:i]) {
				tmp = append(tmp, s[idx:i])
				dfs(i)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0)
	return res
}

func isPalindrome(s string) bool{
	j:=len(s)-1
	for i:=0;i<=len(s)/2;i,j= i+1,j-1{
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
