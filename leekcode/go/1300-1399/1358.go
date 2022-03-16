package _300_1399


/*
1358. 包含所有三种字符的子字符串数目
给你一个字符串 s ，它只包含三种字符 a, b 和 c 。
请你返回 a，b 和 c 都 至少 出现过一次的子字符串数目。
*/
func numberOfSubstrings(s string) int {
	res := 0
	hash := [3]int{}
	hash[0] = 0
	hash[1] = 0
	hash[2] = 0
	l := 0
	r := 0
	for r < len(s) {
		hash[s[r]-'a']++
		for hash[0] > 0 && hash[1] > 0 && hash[2] > 0{
			res += len(s)-r
			hash[s[l]-'a']--
			l++
		}
		r++
	}
	return res
}
