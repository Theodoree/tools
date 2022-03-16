package _100_2199


/*
2168. 每个数字的频率都相同的独特子字符串的数量
给你一个由数字组成的字符串 s，返回 s 中独特子字符串数量，其中的每一个数字出现的频率都相同。
示例1:
输入: s = "1212"
输出: 5
解释: 符合要求的子串有 "1", "2", "12", "21", "1212".
要注意，尽管"12"在s中出现了两次，但在计数的时候只计算一次。
*/
func equalDigitFrequency(s string) int {
	n := len(s)
	hmap := map[string]bool{}
	for k := 1; k <= n; k++ {
		var stat = [10]uint16{}


		for i := 0; i < k; i++ {
			stat[int(s[i]-'0')]++
		}
		if unique(stat) {
			hmap[s[:k]] = true
		}
		for i := 1; i+k <= n; i++ {
			stat[int(s[i-1]-'0')]--
			stat[int(s[i+k-1]-'0')]++
			if unique(stat) {
				hmap[s[i:i+k]] = true
			}
		}
	}
	return len(hmap)
}

func unique(stat [10]uint16) bool {
	i := 0
	for i < len(stat) && stat[i] == 0 {
		i++
	}
	tmp := stat[i]
	i++
	for i < len(stat) {
		if stat[i] > 0 && stat[i] != tmp {
			return false
		}
		i++
	}
	return true
}
