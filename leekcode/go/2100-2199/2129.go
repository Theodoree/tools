package _100_2199

import (
	"reflect"
	"unsafe"
)

/*
2129. 将标题首字母大写
给你一个字符串 title ，它由单个空格连接一个或多个单词组成，每个单词都只包含英文字母。请你按以下规则将每个单词的首字母 大写 ：
如果单词的长度为 1 或者 2 ，所有字母变成小写。
否则，将单词首字母大写，剩余字母变成小写。
请你返回 大写后 的 title 。
*/
func capitalizeTitle(title string) string {
	buf := StringToBytes(title)
	//buf:=[]byte(title)
	var left, right int
	right = findEmpty(buf, left)
	for right <= len(title) {
		cast(buf, left, right)
		left = right + 1
		right = findEmpty(buf, left)
	}

	return title
}
func findEmpty(s []byte, idx int) int {
	for ; idx < len(s); idx++ {
		if s[idx] == ' ' {
			return idx
		}
	}
	return idx
}
func cast(s []byte, left, right int) {
	if right-left > 2 {
		if s[left] >= 'a' {
			s[left] -= 'a' - 'A'
		}
		left++
	}
	for left < right {
		if s[left] <= 'Z' {
			s[left] += 'a' - 'A'
		}
		left++
	}
}

func StringToBytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}
