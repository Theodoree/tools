package _700_1799

/*
1704. 判断字符串的两半是否相似
给你一个偶数长度的字符串 s 。将其拆分成长度相同的两半，前一半为 a ，后一半为 b 。
两个字符串 相似 的前提是它们都含有相同数目的元音（'a'，'e'，'i'，'o'，'u'，'A'，'E'，'I'，'O'，'U'）。注意，s 可能同时含有大写和小写字母。
如果 a 和 b 相似，返回 true ；否则，返回 false 。
*/

func StringToBytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

func halvesAreAlike(s string) bool {

	var count, count1 int
	buf := StringToBytes(s)
	for i := 0; i < len(buf); i++ {
		if buf[i] < 'a' {
			buf[i] += 32
		}
	}
	for i := 0; i < len(s)/2; i++ {

		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}

		switch s[len(s)-i-1] {
		case 'a', 'e', 'i', 'o', 'u':
			count1++
		}

	}

	return count == count1

}
