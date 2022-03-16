package _go

/*
面试题 01.06. 字符串压缩
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
*/
func compressString(S string) string {
	if S == "" {
		return ""
	}

	var buf = make([]byte, len(S)*2)
	var bufIndex = 0
	var last byte
	var count int

	last = S[0]
	count = 1
	for i := 1; i < len(S); i++ {
		if S[i] != last {
			buf[bufIndex] = last
			bufIndex++

			switch {
			case count >= 10000:
				buf[bufIndex] = byte(count/10000 + '0')
				count %= 10000
				bufIndex++
				fallthrough
			case count >= 1000:
				buf[bufIndex] = byte(count/1000 + '0')
				count %= 1000
				bufIndex++
				fallthrough
			case count >= 100:
				buf[bufIndex] = byte(count/100 + '0')
				count %= 100
				bufIndex++
				fallthrough
			case count >= 10:
				buf[bufIndex] = byte(count/10 + '0')
				count %= 10
				bufIndex++
				fallthrough
			default:
				buf[bufIndex] = byte(count + '0')
				bufIndex++
			}

			last = S[i]
			count = 0
		}
		count++
	}

	if count > 0 {
		buf[bufIndex] = last
		bufIndex++
		switch {
		case count >= 10000:
			buf[bufIndex] = byte(count/10000 + '0')
			count %= 10000
			bufIndex++
			fallthrough
		case count >= 1000:
			buf[bufIndex] = byte(count/1000 + '0')
			count %= 1000
			bufIndex++
			fallthrough
		case count >= 100:
			buf[bufIndex] = byte(count/100 + '0')
			count %= 100
			bufIndex++
			fallthrough
		case count >= 10:
			buf[bufIndex] = byte(count/10 + '0')
			count %= 10
			bufIndex++
			fallthrough
		default:
			buf[bufIndex] = byte(count + '0')
			bufIndex++
		}
	}

	if bufIndex >= len(S) {
		return S
	}

	return string(buf[:bufIndex])
}
