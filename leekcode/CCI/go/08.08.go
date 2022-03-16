package _go

/*
面试题 08.08. 有重复字符串的排列组合
有重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合。
*/
func permutation(S string) []string {
	res := []string{
		S[:1],
	}
	var m = make( map[string]struct{})

	for s := 1; s < len(S); s++ {
		t := []string{}

		for _, ss := range res {
			for i := 0; i <= len(ss); i++ {
				str:=ss[:i]+S[s:s+1]+ss[i:]
				if _,ok:=m[str];!ok {
					t = append(t, str)
					m[str] = struct{}{}
				}
			}
		}
		res = t
	}
	return res
}

