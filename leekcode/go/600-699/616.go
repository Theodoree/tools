package _00_699

/*
616. 给字符串添加加粗标签
给你一个字符串 s 和一个字符串列表 words ，你需要将在字符串列表中出现过的 s 的子串添加加粗闭合标签 <b> 和 </b> 。
如果两个子串有重叠部分，你需要把它们一起用一对闭合标签包围起来。同理，如果两个子字符串连续被加粗，那么你也需要把它们合起来用一对加粗标签包围。
返回添加加粗标签后的字符串 s 。
*/

func addBoldTag(s string, dict []string) string {
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
