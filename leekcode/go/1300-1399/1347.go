package _300_1399




/*
1347. 制造字母异位词的最小步骤数
给你两个长度相等的字符串 s 和 t。每一个步骤中，你可以选择将 t 中的 任一字符 替换为 另一个字符。
返回使 t 成为 s 的字母异位词的最小步骤数。
字母异位词 指字母相同，但排列不同（也可能相同）的字符串。
*/
func minSteps(s string, t string) int {
	var buf [2][26]int
	for _,v:=range s{
		buf[0][v-'a']++
	}

	for _,v:=range t{
		buf[1][v-'a']++
	}


	var sum int
	for i:=0;i<26;i++{
		if buf[1][i] < buf[0][i] {
			sum += buf[0][i] - buf[1][i]
		}
	}
	return sum
}
