package _go

import "unsafe"

/*
面试题 08.07. 无重复字符串的排列组合
无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。
*/
var targets = [63]int{
	1, 3, 7, 15,
	31, 63, 127, 255,
	511, 1023, 2047, 4095,
	8191, 16383, 32767, 65535,
	131071, 262143, 524287, 1048575,
	2097151, 4194303, 8388607, 16777215,
	33554431, 67108863, 134217727, 268435455,
	536870911, 1073741823, 2147483647, 4294967295,
	8589934591, 17179869183, 34359738367, 68719476735,
	137438953471, 274877906943, 549755813887, 1099511627775,
	2199023255551, 4398046511103, 8796093022207, 17592186044415,
	35184372088831, 70368744177663, 140737488355327, 281474976710655,
	562949953421311, 1125899906842623, 2251799813685247, 4503599627370495,
	9007199254740991, 18014398509481983, 36028797018963967, 72057594037927935,
	144115188075855871, 288230376151711743, 576460752303423487, 1152921504606846975,
	2305843009213693951, 4611686018427387903, 9223372036854775807,
}

func permutation(S string) []string {
	var result []string
	target := targets[len(S)-1]
	backtracking(S, 0, target, nil, &result)

	return result
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func backtracking(s string, idxFlag int, target int, cur []byte, result *[]string) {
	if target == idxFlag {
		*result = append(*result, BytesToString(cur))
		return
	}

	for i := 0; i < len(s); i++ {
		if idxFlag&(1<<i) > 0 {
			continue
		}

		buf := make([]byte, len(cur))
		copy(buf, cur)
		buf = append(buf, s[i])

		backtracking(s, idxFlag|1<<i, target, buf, result)
	}

}

func permutation(S string) []string {
	res := []string{
		S[:1],
	}

	for s := 1; s < len(S); s++ {
		t := []string{}

		for _, ss := range res {
			for i := 0; i <= len(ss); i++ {
				t = append(t, ss[:i]+S[s:s+1]+ss[i:])
			}
		}
		res = t
	}
	return res
}
