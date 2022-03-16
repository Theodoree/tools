package _000_1099

/*
1002. 查找常用字符
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。

你可以按任意顺序返回答案。

*/
func commonChars(A []string) []string {

	var filter [26]int
	filter[0] = -1

	for i := 0; i < len(A); i++ {
		var f [26]int
		for j := 0; j < len(A[i]); j++ {
			f[A[i][j]-'a']++
		}

		if filter[0] == -1{
			filter = f
		}

		for i:=0;i<26;i++{
			if f[i] < filter[i]{
				filter[i] = f[i]
			}
		}
	}

	var result []string
	var n  int
	for k,v:=range filter{
		for i:=0;i<v;i++{
			n = k+'a'
			result = append(result,string(n))
		}

	}
	return result
}

