package _00_299


/*
249. 移位字符串分组
给定一个字符串，对该字符串可以进行 “移位” 的操作，也就是将字符串中每个字母都变为其在字母表中后续的字母，比如："abc" -> "bcd"。这样，我们可以持续进行 “移位” 操作，从而生成如下移位序列：
"abc" -> "bcd" -> ... -> "xyz"
给定一个包含仅小写字母字符串的列表，将该列表中所有满足 “移位” 操作规律的组合进行分组并返回。
*/

func groupStrings(strings []string) (ans [][]string) {
	patternStrsRel:=map[string][]string{}
	for _,str:=range strings{
		newstr:=""
		for i:=range str{
			newstr+=string((str[i]-str[0]+26)%26+'a')
		}
		patternStrsRel[newstr]=append(patternStrsRel[newstr],str)
	}

	fmt.Println(patternStrsRel)
	for _,v:=range patternStrsRel{
		ans=append(ans,v)
	}

	return
}
