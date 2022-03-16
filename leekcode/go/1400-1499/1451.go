package _400_1499

/*
1451. 重新排列句子中的单词
「句子」是一个用空格分隔单词的字符串。给你一个满足下述格式的句子 text :
句子的首字母大写
text 中的每个单词都用单个空格分隔。
请你重新排列 text 中的单词，使所有单词按其长度的升序排列。如果两个单词的长度相同，则保留其在原句子中的相对顺序。
请同样按上述格式返回新的句子。
*/
type item struct {
	str string
	pos int
}

func arrangeWords(text string) string {
	array := make([]item, 0, len(text)/100)
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			continue
		}
		for j := i; j < len(text); j++ {
			if text[j] == ' ' {
				array = append(array, item{text[i:j], len(array)})
				i = j
				break
			} else if j == len(text)-1 {
				array = append(array, item{text[i:], len(array)})
				i = j
				break
			}
		}
	}
	sort.Sort(stringLen(array))
	result := make([]byte, len(text))
	copy(result, array[0].str)
	if array[0].str[0] >= 'a' && array[0].str[0] <= 'z' {
		result[0] = result[0] - 'a' + 'A'
	}
	l := len(array[0].str)
	for i := 1; i < len(array); i++ {
		result[l] = ' '
		l++
		copy(result[l:], array[i].str)
		if array[i].str[0] >= 'A' && array[i].str[0] <= 'Z' {
			result[l] = result[l] - 'A' + 'a'
		}
		l += len(array[i].str)
	}
	return string(result)
}

type stringLen []item

func (sl stringLen) Len() int {
	return len(sl)
}

func (sl stringLen) Less(i, j int) bool {
	if len(sl[i].str) == len(sl[j].str) {
		return sl[i].pos < sl[j].pos
	}
	return len(sl[i].str) < len(sl[j].str)
}

func (sl stringLen) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}
