package _000_2099

/*
2042. 检查句子中的数字是否递增
句子是由若干 token 组成的一个列表，token 间用 单个 空格分隔，句子没有前导或尾随空格。每个 token 要么是一个由数字 0-9 组成的不含前导零的 正整数 ，要么是一个由小写英文字母组成的 单词 。

示例，"a puppy has 2 eyes 4 legs" 是一个由 7 个 token 组成的句子："2" 和 "4" 是数字，其他像 "puppy" 这样的 tokens 属于单词。
给你一个表示句子的字符串 s ，你需要检查 s 中的 全部 数字是否从左到右严格递增（即，除了最后一个数字，s 中的 每个 数字都严格小于它 右侧 的数字）。

如果满足题目要求，返回 true ，否则，返回 false 。
*/

const f = '0'

func areNumbersAscending(s string) bool {
	var num [2]int
	num[1] = -1
	for i := 0; i < len(s); i++ {

		switch {
		case s[i] >= '0' && s[i] <= '9':
			var _num int
			var flag bool
			for i < len(s) && s[i] != ' ' {
				_num = _num*10 + int(s[i]-'0')
				i++
			}
			if flag {
				continue
			}
			num[0] = num[1]
			num[1] = _num
			if num[0] >= num[1] {
				return false
			}

		default:
			continue

		}

	}

	return true
}
