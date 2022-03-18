package _00_399

/*
385. 迷你语法分析器
给定一个字符串 s 表示一个整数嵌套列表，实现一个解析它的语法分析器并返回解析的结果 NestedInteger 。
列表中的每个元素只可能是整数或整数嵌套列表
*/
func deserialize(s string) *NestedInteger {

	var arr []*NestedInteger

	var result *NestedInteger
	const (
		leftToken  = '['
		rightToken = ']'
		pointToken = ','
	)

	hasNum := false
	sign := false
	var num int
	for len(s) > 0 {
		switch s[0] {
		case leftToken:
			cur := NestedInteger{}
			arr = append(arr, &cur)
		case rightToken:
			if hasNum {
				cur := NestedInteger{}
				cur.SetInteger(num)
				last := arr[len(arr)-1]
				last.Add(cur)
				num = 0
			}
			result = arr[len(arr)-1]
			if len(arr)-2 >= 0 {
				arr[len(arr)-2].Add(*result)
			}
			arr = arr[:len(arr)-1]
			sign = false
			hasNum = false
		case pointToken:
			if hasNum {
				cur := NestedInteger{}
				cur.SetInteger(num)
				last := arr[len(arr)-1]
				last.Add(cur)
				num = 0
			}
			sign = false
			hasNum = false
		case '-':
			sign = true
		default:
			hasNum = true
			if sign {
				num = num*10 - int(s[0]-'0')
			} else {
				num = num*10 + int(s[0]-'0')
			}
		}
		s = s[1:]
	}

	if hasNum {
		cur := &NestedInteger{}
		cur.SetInteger(num)
		return cur
	}

	return result
}
