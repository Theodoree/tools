package _go

/*
面试题 01.09. 字符串轮转
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。

示例1:

 输入：s1 = "waterbottle", s2 = "erbottlewat"
 输出：True
示例2:

 输入：s1 = "aa", s2 = "aba"
 输出：False
提示：

字符串长度在[0, 100000]范围内。
说明:

*/

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	for i := 0; i < len(s2); i++ {
		s1Index := 0
		s2Index := i
		for s1[s1Index] == s2[s2Index] {
			s1Index++
			s2Index++

			if s2Index == len(s2) {
				s2Index = 0
			}
			if s1Index == len(s1) {
				return true
			}

		}

	}

	return false
}
