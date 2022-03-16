package _00_699

/*
696. 计数二进制子串
给定一个字符串 s，统计并返回具有相同数量 0 和 1 的非空（连续）子字符串的数量，并且这些子字符串中的所有 0 和所有 1 都是成组连续的。
重复出现（不同位置）的子串也要统计它们出现的次数。
*/
func countBinarySubstrings(s string) int {
	var pre, curr, count int
	curr = 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			curr++
		} else {
			pre = curr
			curr = 1
		}
		if pre >= curr {
			// 本身只有0和1构造,所以先计算某一种,再计算另外一种,当pre大于curr的时候,
			// 例子:
			// 00001 取01
			// 000011 取0011
			// 0000111 取0s00111
			count++
		}
	}
	return count

}
