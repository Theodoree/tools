package _go




/*
面试题 10.02. 变位词组
编写一种方法，对字符串数组进行排序，将所有变位词组合在一起。变位词是指字母相同，但排列不同的字符串。
注意：本题相对原题稍作修改
*/
func groupAnagrams(strs []string) [][]string {

	var m = make(map[[26]uint16][]string)
	for _,str:=range strs{
		buf:=[26]uint16{}
		for _,v:=range str{
			buf[v-'a']++
		}

		m[buf]= append(m[buf],str)
	}

	var result [][]string
	for _,v:=range m{
		result = append(result,v)
	}
	return result

}

