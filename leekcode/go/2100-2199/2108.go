package _100_2199

/*
2108. 找出数组中的第一个回文字符串
给你一个字符串数组 words ，找出并返回数组中的 第一个回文字符串 。如果不存在满足要求的字符串，返回一个 空字符串 "" 。
回文字符串 的定义为：如果一个字符串正着读和反着读一样，那么该字符串就是一个 回文字符串 。
*/

func firstPalindrome(words []string) string {

	for i := 0; i < len(words); i++ {
		str := words[i]
		left := len(str) / 2
		right := left
		if len(str)%2 == 0 {
			left--
		}

		for left >= 0 && str[left] == str[right] {
			left--
			right++
		}
		if left == -1 {
			return str
		}
	}

	return ""
}
