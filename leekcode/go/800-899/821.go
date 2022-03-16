package _00_899

import "math"

/*
821. 字符的最短距离

给定一个字符串 S 和一个字符 C。返回一个代表字符串 S 中每个字符到字符串 S 中的字符 C 的最短距离的数组。

示例 1:

输入: S = "loveleetcode", C = 'e'
输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
说明:

字符串 S 的长度范围为 [1, 10000]。
C 是一个单字符，且保证是字符串 S 里的字符。
S 和 C 中的所有字母均为小写字母。
*/
func shortestToChar(S string, C byte) []int {

	var c []int
	for k, _ := range S {

		if S[k] == C {
			c = append(c, k)
		}
	}
	var result = make([]int, len(S), len(S))
	var left, right int
	var index int
	if len(c) > 0 {
		left = c[0]
	}

	for i := 0; i < len(S); i++ {
		if i >= right && index+1 < len(c) {
			if right != 0 {
				left = right
			}
			index++
			right = c[index]
		}

		if i < c[0] {
			result[i] = c[0] - i
		} else if i != left && i != right {
			if i-left > int(math.Abs(float64(i-right))) {
				result[i] = int(math.Abs(float64(i - right)))
			} else {
				result [i] = i - left

			}
		}

	}

	return result
}