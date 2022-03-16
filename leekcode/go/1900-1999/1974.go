package _900_1999

import "math"

/*
1974. 使用特殊打字机键入单词的最少时间
有一个特殊打字机，它由一个圆盘和一个指针组成，圆盘上标有小写英文字母'a'到'z'。
只有当指针指向某个字母时，它才能被键入。指针初始时指向字符 'a' 。
*/

func minTimeToType(word string) int {
	res := 0.0
	start := 'a'
	for _, v := range word {
		dis := math.Abs(float64(v - start))
		res += math.Min(dis, 26-dis)
		res++
		start = v
	}
	return int(res)

}
