package _00_799

func boldWords(dict []string, s string) string {
	mask := make([]bool, len(s))
	end := 0
	for i := 0; i < len(s); i++ {
		for _, word := range dict { //遍历每个字符，进入dict的就打mark
			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
				end = max(end, i+len(word)) //end 要一直取最大值 ,Input: s = "aaabbcc", words = ["aaa","aab","bc"] Output: "<b>aaabbc</b>c"
			}
		}
		mask[i] = end > i //end 多走了一步 ， 相当于end-1>=i
	}

	var ans strings.Builder
	for i := 0; i < len(s); i++ {
		if mask[i] && (i == 0 || !mask[i-1]) {
			ans.WriteString("<b>")
		}
		ans.WriteByte(s[i])
		if mask[i] && (i == len(s)-1 || !mask[i+1]) {
			ans.WriteString("</b>")
		}
	}

	return ans.String()
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
