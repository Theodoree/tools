package _00_799





/*
763. 划分字母区间
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
*/
func partitionLabels(s string) []int {
	var buf [26]int
	for idx,v:=range s{
		buf[v-'a'] = idx
	}
	max:= func(i,j int) int {
		if i > j {
			return i
		}
		return j
	}
	var end int
	var start int
	var result []int
	for i:=0;i<len(s);i++{

		end = max(end,buf[s[i]-'a'])
		if i== end {
			result=append(result,end-start+1)
			start = i+1
		}
	}
	return result

}

