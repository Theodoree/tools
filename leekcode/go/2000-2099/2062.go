package _000_2099

/*
2062. 统计字符串中的元音子字符串
子字符串 是字符串中的一个连续（非空）的字符序列。
元音子字符串 是 仅 由元音（'a'、'e'、'i'、'o' 和 'u'）组成的一个子字符串，且必须包含 全部五种 元音。
给你一个字符串 word ，统计并返回 word 中 元音子字符串的数目 。
*/
func countVowelSubstrings(word string) int {
	var count int
	for i := 0; i < len(word); i++ {
		var a int
	loop:
		for j := i; j < len(word); j++ {
			switch word[j] {
			case 'a':
				a |= 1 << 0
			case 'e':
				a |= 1 << 1
			case 'i':
				a |= 1 << 2
			case 'o':
				a |= 1 << 3
			case 'u':
				a |= 1 << 4
			default:
				break loop
			}

			if a == 31 {
				count++
			}
		}
	}

	return count
}
