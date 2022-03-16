package _400_1499

import "unsafe"

/*
1417. 重新格式化字符串
给你一个混合了数字和字母的字符串 s，其中的字母均为小写英文字母。
请你将该字符串重新格式化，使得任意两个相邻字符的类型都不同。也就是说，字母后面应该跟着数字，而数字后面应该跟着字母。
请你返回 重新格式化后 的字符串；如果无法按要求重新格式化，则返回一个 空字符串 。
*/
func reformat(s string) string {
	if len(s) <= 1 {
		return s
	}
	num := make([]byte, 0, len(s))
	alphabet := make([]byte, 0, len(s))

	for _, v := range []byte(s) {
		if v <= '9' {
			num = append(num, v)
		} else {
			alphabet = append(alphabet, v)
		}
	}

	if len(num) == 0 || len(alphabet) == 0 {
		return ""
	}
	if len(num) < len(alphabet) {
		num, alphabet = alphabet, num
	}

	if len(num)-len(alphabet) > 1 {
		return ""
	}

	result := []byte(s)
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			result[i] = num[0]
			num = num[1:]
		} else {
			result[i] = alphabet[0]
			alphabet = alphabet[1:]
		}
	}

	return BytesToString(result)
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
