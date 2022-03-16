package _go

/*
剑指 Offer II 019. 最多删除一个字符得到回文
给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。
*/

func validPalindrome(s string) bool {

	var left, right int

	sByte := []byte(s)
	right = len(s) - 1

	for left < right {
		if sByte[left] != sByte[right] {
			return ValidPalindrome(sByte, left, right-1) || ValidPalindrome(sByte, left+1, right)
		}
		left++
		right--
	}

	return true

}

func ValidPalindrome(chars []byte, l, r int) bool {
	for l < r {
		if chars[l] != chars[r] {
			return false
		}
		l++
		r--
	}
	return true
}
