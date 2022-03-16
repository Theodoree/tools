package _100_1199

/*
1100. 长度为 K 的无重复字符子串
给你一个字符串 S，找出所有长度为 K 且不含重复字符的子串，请你返回全部满足要求的子串的 数目。
*/
func numKLenSubstrNoRepeats(s string, k int) int {

	var i,j int
	var arr [26]byte
	var result int

	for j < len(s){
		for  arr[s[j]-'a'] > 0 { // 那么这个元素已经出现过了,left往前滚
			arr[s[i]-'a']--
			i++
		}

		arr[s[j]-'a']++

		if j-i+1 == k {
			result ++

			arr[s[i] - 'a']--
			i++
		}
		j++
	}

	return result
}


