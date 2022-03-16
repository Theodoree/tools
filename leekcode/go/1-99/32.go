package __99




/*
32. 最长有效括号
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
*/
func longestValidParentheses(s string) int {
	dp:=make([]int,len(s))
	max:= func(i,j int) int {
		if i > j {
			return  i
		}
		return j
	}
	var maxans int
	for i:=1;i<len(s);i++{
		if s[i] == ')' {
			if s[i-1] == '('{ // 两种情况 ()  ()()
				if i >= 2  { // ()()
					dp[i] = dp[i-2]+2
				}else{ // ()
					dp[i] = 2
				}
			}else if  i - dp[i-1] >0 && s[i-dp[i-1]-1] == '('  {
				if i-dp[i-1] >=2 { // ()((()))
					dp[i] = dp[i-1] + dp[i-dp[i-1] - 2] + 2
				}else{ // ()()
					dp[i] = dp[i-1]+2
				}
			}
			maxans = max(dp[i],maxans)
		}


	}

	return maxans

}
