package _00_399



/*
320. 列举单词的全部缩写
单词的 广义缩写词 可以通过下述步骤构造：先取任意数量的 不重叠、不相邻 的子字符串，再用它们各自的长度进行替换。
例如，"abcde" 可以缩写为：
"a3e"（"bcd" 变为 "3" ）
"1bcd1"（"a" 和 "e" 都变为 "1"）
"5" ("abcde" 变为 "5")
"abcde" (没有子字符串被代替)
然而，这些缩写是 无效的 ：
"23"（"ab" 变为 "2" ，"cde" 变为 "3" ）是无效的，因为被选择的字符串是相邻的
"22de" ("ab" 变为 "2" ， "bc" 变为 "2")  是无效的，因为被选择的字符串是重叠的
给你一个字符串 word ，返回 一个由 word 的所有可能 广义缩写词 组成的列表 。按 任意顺序 返回答案。
*/

var karr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14"}

func generateAbbreviations(word string) []string {
	var arr []string
	Backtracking(word, 0, "", 0, &arr)
	return arr
}

func Backtracking(word string, idx int, cur string, k int, arr *[]string) {
	if idx == len(word) {
		if k != 0 {
			cur += karr[k]
		}
		*arr = append(*arr, cur)
		return
	}

	if k == 0 { // 说明前面是数字
		Backtracking(word, idx+1, cur+string(word[idx]), 0, arr)
	} else {
		Backtracking(word, idx+1, cur+karr[k]+string(word[idx]), 0, arr)
	}
	Backtracking(word, idx+1, cur, k+1, arr)

}

