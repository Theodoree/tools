package _700_1799

import "unsafe"

/*
1790. 仅执行一次字符串交换能否使两个字符串相等
给你长度相等的两个字符串 s1 和 s2 。一次 字符串交换 操作的步骤如下：选出某个字符串中的两个下标（不必不同），并交换这两个下标所对应的字符。
如果对 其中一个字符串 执行 最多一次字符串交换 就可以使两个字符串相等，返回 true ；否则，返回 false 。
*/
func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	// 以s1为准
	l1 := -1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if l1 == -1 {
				l1 = i
			} else {
				b := unsafeToBuf(s2)
				b[l1], b[i] = b[i], b[l1]
				return s1 == s2
			}
		}
	}

	return false

}
func unsafeToBuf(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.StringHeader)(unsafe.Pointer(&s)).Data,
		Len:  len(s),
		Cap:  len(s),
	}))

}
