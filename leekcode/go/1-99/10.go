package __99

import (
	"fmt"
	"regexp"
)

/*
剑指 Offer 19. 正则表达式匹配
请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母以及字符 . 和 *，无连续的 '*'。
*/

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	t := make([][]bool, m+1)
	t[0] = make([]bool, n+1)
	t[0][0] = true
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			t[0][j] = t[0][j-2]
		}
	}
	for i := 1; i <= m; i++ {
		t[i] = make([]bool, n+1)
		if 0 != n {
			t[i][1] = t[i-1][0] && (s[i-1] == p[0] || p[0] == '.')
		}
		for j := 2; j <= n; j++ {
			if p[j-1] == '*' {
				t[i][j] = t[i][j-2] || ((s[i-1] == p[j-2] || p[j-2] == '.') && t[i-1][j])
			} else {
				t[i][j] = t[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	fmt.Println(t)
	return t[m][n]
}


func isMatch(s string, p string) bool {
	ok, err := regexp.MatchString("^"+p+"$",s)
	if err != nil{
		return false
	}
	return ok
}


func main() {

	fmt.Println(isMatch(`aab`, `c*a*b`))

}
