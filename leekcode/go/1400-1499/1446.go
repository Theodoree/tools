package _400_1499

/*
1446. 连续字符
给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。
请你返回字符串的能量。
*/
func maxPower(s string) int {
	var maxCount int

	count := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			count++
			continue
		}

		if count > maxCount {
			maxCount = count
		}
		count = 1
	}
	if count > maxCount {
		maxCount = count
	}
	return maxCount
}
