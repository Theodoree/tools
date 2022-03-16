package _800_1899

/*
1805. 字符串中不同整数的数目
给你一个字符串 word ，该字符串由数字和小写英文字母组成。
请你用空格替换每个不是数字的字符。例如，"a123bc34d8ef34" 将会变成 " 123  34 8  34" 。注意，剩下的这些整数为（相邻彼此至少有一个空格隔开）："123"、"34"、"8" 和 "34" 。
返回对 word 完成替换后形成的 不同 整数的数目。
只有当两个整数的 不含前导零 的十进制表示不同， 才认为这两个整数也不同。
*/
func numDifferentIntegers(word string) int {
	var m = make(map[int]struct{}, 3)
	var m1 = make(map[string]struct{}, 0)
	var num, start int
	var overflow bool

	start = -1

	for i := 0; i < len(word); i++ {
		if !(word[i] >= '0' && word[i] <= '9') {
			if start >= 0 {
				if overflow {
					m1[word[start:i]] = struct{}{}
				} else {
					m[num] = struct{}{}
				}
				num = 0
				start = -1
				overflow = false
			}
			continue
		}

		if start == -1 {
			start = i
		}
		old := num
		num = num*10 + int(word[i]) - '0'
		if num < old {
			overflow = true
		}
	}

	if start >= 0 {
		if overflow {
			m1[word[start:]] = struct{}{}
		} else {
			m[num] = struct{}{}
		}
		num = 0
		start = -1
		overflow = false
	}

	return len(m) + len(m1)
}
