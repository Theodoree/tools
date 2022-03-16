package _go


/*
面试题 16.20. T9键盘
在老式手机上，用户通过数字键盘输入，手机将提供与这些数字相匹配的单词列表。每个数字映射到0至4个字母。给定一个数字序列，实现一个算法来返回匹配单词的列表。你会得到一张含有有效单词的列表。映射如下图所示：
*/
func getValidT9Words(num string, words []string) []string {
	var result []string
	for _,word:=range words{
		ok:=true
		for idx:=range word{
			left:=byte(0)
			right:=byte(0)
			_num:=num[idx]-'0'
			switch _num {
			case 7:
				left = 'p'
				right = 's'
			case 8:
				left = 't'
				right = 'v'
			case 9:
				left = 'w'
				right = 'z'
			default:
				left = (_num-2)*3+'a'
				right =(_num-1)*3+'a'-1
			}

			if word[idx] < left ||  word[idx] > right{
				ok = false
				break
			}
		}
		if ok {
			result = append(result,word)
		}
	}

	return result
}

