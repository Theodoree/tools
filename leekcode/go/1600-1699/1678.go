package _600_1699

import "unsafe"

/*
1678. 设计 Goal 解析器
请你设计一个可以解释字符串 command 的 Goal 解析器 。command 由 "G"、"()" 和/或 "(al)" 按某种顺序组成。
Goal 解析器会将 "G" 解释为字符串 "G"、"()" 解释为字符串 "o" ，"(al)" 解释为字符串 "al" 。然后，按原顺序将经解释得到的字符串连接成一个字符串。
给你字符串 command ，返回 Goal 解析器 对 command 的解释结果。
*/
func interpret(command string) string {
	var buf = make([]byte, 0, len(command))
	for i := 0; i < len(command); i++ {

		switch command[i] {
		case 'G':
			buf = append(buf, 'G')
		case '(':
			if command[i+1] == ')' {
				i++
				buf = append(buf, 'o')
				continue
			}
			buf = append(buf, 'a', 'l')
			i += 2
		}

	}
	return BytesToString(buf)
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
