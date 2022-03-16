package _00_199


/*
159. 至多包含两个不同字符的最长子串
给定一个字符串 s ，找出 至多 包含两个不同字符的最长子串 t ，并返回该子串的长度。
*/
func lengthOfLongestSubstringTwoDistinct(s string) int {
	var max int
	var buf [math.MaxUint8]bool
	var idx = make([]int,0,2)
	for i:=0;i<len(s);i++{
		right:=i
		for j:=i;j<len(s);j++{
			val:=buf[s[j]]
			if len(idx)>=2  && !val {
				break
			}
			buf[s[j]] = true
			right = j
			if !val {
				idx = append(idx,j)
			}
		}

		if right-i+1 > max {
			max =right-i+1
		}

		min:=math.MaxInt64
		for _,v:=range idx {
			if v < min { // 选择最小的
				min = v
			}
			buf[s[v]] =false
		}

		if len(idx) > 0 && min > i { // 如果大于i,那么就skip这个字符
			i = min
		}
		idx = idx[:0]
	}

	return  max
}


