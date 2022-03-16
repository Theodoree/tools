package _00_399



/*
340. 至多包含 K 个不同字符的最长子串
给定一个字符串 s ，找出 至多 包含 k 个不同字符的最长子串 T。
*/
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	res, diff := 0, 0
	n := len(s)
	l, r := 0, -1
	hash := [128]int{}
	for l < n {
		if r + 1 < n && ((hash[s[r+1]] == 0 && diff+1 <= k) ||  hash[s[r+1]] != 0) {
			r++
			if hash[s[r]] == 0 {
				diff++
			}
			hash[s[r]]++
		} else {
			hash[s[l]]--
			if hash[s[l]] == 0 {
				diff--
			}
			l++
		}

		if r - l + 1 > res {
			res = r -l + 1
		}

	}

	return res
}
