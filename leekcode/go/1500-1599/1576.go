package _500_1599

import (
	"reflect"
	"unsafe"
)

/*
1576. 替换所有的问号
给你一个仅包含小写英文字母和 '?' 字符的字符串 s，请你将所有的 '?' 转换为若干小写字母，使最终的字符串不包含任何 连续重复 的字符。

注意：你 不能 修改非 '?' 字符。

题目测试用例保证 除 '?' 字符 之外，不存在连续重复的字符。

在完成所有转换（可能无需转换）后返回最终的字符串。如果有多个解决方案，请返回其中任何一个。可以证明，在给定的约束条件下，答案总是存在的。
*/

func StringToBytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func modifyString(s string) string {

	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 && s[0] == '?' {
		return "a"
	}

	bufArr := StringToBytes(s)
	bytez := byte('a')

	if len(bufArr) > 0 && bufArr[0] == '?' {
		if bufArr[1] != '?' {
			if bufArr[1]%'a' > 0 {
				bufArr[0] = bufArr[1] - 1
			} else {
				bufArr[0] = bufArr[1] + 1
			}
		} else {
			bufArr[0] = 'a'
		}
	}

	for i := 1; i < len(bufArr); i++ {
		if bufArr[i] == '?' {
			for bytez == bufArr[i-1] || (i+1 < len(bufArr) && bytez == bufArr[i+1]) {
				bytez++
				if bytez > 'z' {
					bytez = 'a'
				}
			}
			bufArr[i] = bytez
		}
	}

	return BytesToString(bufArr)
}
