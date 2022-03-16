package _100_1199



/*
1143. 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	var max = func(i,j int) int {
		if i > j {
			return i
		}
		return j
	}
	dp:=make([][]int,len(text1)+1)
	for  idx:=range dp {
		dp[idx] = make([]int,len(text2)+1)
	}

	for i:=1;i<len(dp);i++{

		for j:=1;j<len(dp[i]);j++{
			if text1[i-1] == text2[j-1]{
				dp[i][j] = dp[i-1][j-1] + 1
			}else{
				dp[i][j] =
					max(dp[i][j-1],
						dp[i-1][j])
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}