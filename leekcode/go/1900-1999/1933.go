package _900_1999

/*
1933. 判断字符串是否可分解为值均等的子串
一个字符串的所有字符都是一样的，被称作等值字符串。

举例，"1111" 和 "33" 就是等值字符串。
相比之下，"123"就不是等值字符串。
规则：给出一个数字字符串s，将字符串分解成一些等值字符串，如果有且仅有一个等值子字符串长度为2，其他的等值子字符串的长度都是3.

如果能够按照上面的规则分解字符串s，就返回真，否则返回假。

子串就是原字符串中连续的字符序列。
*/
func isDecomposable(s string) bool {

	var flag bool
	count := 1
	lastByte := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == lastByte {
			count++
		} else {
			if count == 5 || count == 2 || count%3 == 2 {
				if flag {
					return false
				}
				flag = true
			} else if count%3 != 0 {
				return false
			}
			count = 1
			lastByte = s[i]
		}
	}
	switch count {
	case 5, 2:
		if flag {
			return false
		}
		return true
	default:
		if !flag && count%3 == 2 {
			return true
		}

		if count%3 != 0 {
			return false
		}
	}

	return flag
}
