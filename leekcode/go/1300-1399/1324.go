package _300_1399

import "strings"

/*
1324. 竖直打印单词
给你一个字符串 s。请你按照单词在 s 中的出现顺序将它们全部竖直返回。
单词应该以字符串列表的形式返回，必要时用空格补位，但输出尾部的空格需要删除（不允许尾随空格）。
每个单词只能放在一列上，每一列中也只能有一个单词。
*/
func printVertically(s string) []string {
	words:=strings.Split(s," ")

	var max int
	for _,v:=range words{
		if len(v) > max {
			max = len(v)
		}
	}

	arr:=make([]string,0,max)
	for i:=0;i<max;i++{
		var b strings.Builder
		b.Grow(len(words))
		for j:=0;j<len(words);j++{
			if len(words[j]) <= i {
				b.WriteByte(' ')
			}else{
				b.WriteByte(words[j][i])
			}
		}
		str:=b.String()
		for str[len(str)-1] == ' '{
			str = str[:len(str)-1]
		}
		arr = append(arr,str)
	}

	return arr
}
