package _100_1199

/*
1180. 统计只含单一字母的子串
给你一个字符串 S，返回只含 单一字母 的子串个数。

aaaa
aaaa 1
aaa  2
aa   3
a    4
sum = n * (n+1) /2
sum = 4 * 5 / 2  = 1 + 2 + 3 + 4


*/
func countLetters(s string) int {
	var sum int
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			count++
		} else {
			sum += count * (count + 1) / 2
			count = 1
		}
	}
	sum += count * (count + 1) / 2

	return sum
}
