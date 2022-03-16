package _go


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
