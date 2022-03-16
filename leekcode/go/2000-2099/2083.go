package _000_2099



/*
2083. 求以相同字母开头和结尾的子串总数
给你一个仅由小写英文字母组成的，  下标从 0 开始的字符串 s 。返回 s 中以相同字符开头和结尾的子字符串总数。
子字符串是字符串中连续的非空字符序列。
*/
func numberOfSubstrings(s string) int64 {
	var buf [26]int
	for _,v:=range s{
		buf[v-'a']++
	}

	var sum int
	for _,v:=range buf {
		if v > 0 {
			sum+=(1+v)*v/2
		}
	}
	return int64(sum)
}
