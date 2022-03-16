package _go

/*
面试题 01.05. 一次编辑
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
*/
func oneEditAway(first string, second string) bool {

	if math.Abs(float64(len(first)-len(second))) >= 2 {
		return false
	}

	var action int
	if len(first) > len(second) { // 插入
		action = 1
	} else if len(first) < len(second) { // 删除
		action = 2
	} else { // 替换
		action = 3
	}

	switch action {
	case 1:
		for i := 0; i < len(second); i++ {
			if second[i] == first[i] {
				continue
			}
			if second[i:] == first[i+1:] {
				return true
			} else {
				return false
			}
		}
	case 2:
		for i := 0; i < len(first); i++ {
			if first[i] == second[i] {
				continue
			}
			if first[i:] == second[i+1:] {
				return true
			} else {
				return false
			}
		}
		return true
	case 3:
		var flag bool
		for i := 0; i < len(second); i++ {
			if second[i] == first[i] {
				continue
			}
			if flag {
				return false
			}
			flag = true
		}
		return true
	}

	return true
}
