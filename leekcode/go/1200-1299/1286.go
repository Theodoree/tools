package _200_1299

import "unsafe"

/*
1286. 字母组合迭代器
请你设计一个迭代器类，包括以下内容：
一个构造函数，输入参数包括：一个 有序且字符唯一 的字符串 characters（该字符串只包含小写英文字母）和一个数字 combinationLength 。
函数 next() ，按 字典序 返回长度为 combinationLength 的下一个字母组合。
函数 hasNext() ，只有存在长度为 combinationLength 的下一个字母组合时，才返回 True；否则，返回 False。
*/

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type CombinationIterator struct {
	arr []string
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	var c CombinationIterator
	backtracking(characters, 0, combinationLength, nil, &c.arr)
	return c
}

func (this *CombinationIterator) Next() string {
	if !this.HasNext() {
		return ""
	}
	str := this.arr[0]
	this.arr = this.arr[1:]
	return str

}
func (this *CombinationIterator) HasNext() bool {
	return len(this.arr) > 0
}

func backtracking(characters string, idx int, maxLen int, cur []byte, arr *[]string) {
	if len(cur) == maxLen {
		*arr = append(*arr, BytesToString(cur))
		return
	}
	for i := idx; i < len(characters); i++ {
		buf := make([]byte, len(cur))
		copy(buf, cur)
		buf = append(buf, characters[i])
		backtracking(characters, i+1, maxLen, buf, arr)
	}

}
