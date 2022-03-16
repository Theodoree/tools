package _go

import (
	"sort"
	"unsafe"
)

/*
剑指 Offer 38. 字符串的排列
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
*/
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func permutation(s string) []string {
	var result []string
	var target int64
	for i := 0; i < len(s); i++ {
		target |= 1 << i
	}
	Backtracking(s, 0, target, nil, &result)

	var m = map[string]struct{}{}
	for _, v := range result {
		m[v] = struct{}{}
	}

	result = result[:0]
	for k := range m {
		result = append(result, k)
	}

	return result
}

func Backtracking(s string, flag int64, target int64, cur []byte, result *[]string) {
	if flag == target {
		*result = append(*result, BytesToString(cur))
		return
	}

	for i := 0; i < len(s); i++ {
		if flag&(1<<i) > 0 {
			continue
		}
		buf := make([]byte, len(cur), len(cur)+1)
		copy(buf, cur)
		buf = append(buf, s[i])
		Backtracking(s, flag|1<<i, target, buf, result)
	}

}

func permutation(s string) []string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	res := []string{string(bytes)}
	if len(bytes) < 2 {
		return res
	}

	var reverse = func(begin, end int) {
		for begin < end {
			bytes[begin], bytes[end] = bytes[end], bytes[begin]
			begin++
			end--
		}
	}

	for {
		i := len(bytes) - 1
		for i > 0 && bytes[i] <= bytes[i-1] {
			i--
		}
		if i == 0 {
			return res
		}

		j := len(bytes) - 1
		for j >= 0 && bytes[i-1] >= bytes[j] {
			j--
		}
		bytes[j], bytes[i-1] = bytes[i-1], bytes[j]
		reverse(i, len(bytes)-1)
		res = append(res, string(bytes))
	}
	return res
}
