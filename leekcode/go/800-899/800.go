package _00_899

import "fmt"

/*
800. 相似 RGB 颜色
RGB 颜色 "#AABBCC" 可以简写成 "#ABC" 。
例如，"#15c" 其实是 "#1155cc" 的简写。
现在，假如我们分别定义两个颜色 "#ABCDEF" 和 "#UVWXYZ"，则他们的相似度可以通过这个表达式 -(AB - UV)^2 - (CD - WX)^2 - (EF - YZ)^2 来计算。
那么给你一个按 "#ABCDEF" 形式定义的字符串 color 表示 RGB 颜色，请你以字符串形式，返回一个与它相似度最大且可以简写的颜色。（比如，可以表示成类似 "#XYZ" 的形式）
任何 具有相同的（最大）相似度的答案都会被视为正确答案。
*/
func similarRGB(color string) string {
	return fmt.Sprintf("#%s%s%s", format(color[1:3]), format(color[3:5]), format(color[5:]))
}

func format(comp string) string {
	var first, second uint8
	if comp[0] >= 'a' {
		first = comp[0] - 'a' + 10
	} else {
		first = comp[0] - '0'
	}
	if comp[1] >= 'a' {
		second = comp[1] - 'a' + 10
	} else {
		second = comp[1] - '0'
	}
	value := first*16 + second
	q := value / 17
	if value%17 > 8 {
		q += 1
	}
	return fmt.Sprintf("%02x", 17*q)
}
