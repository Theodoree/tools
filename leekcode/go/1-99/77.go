package __99

import (
	"strconv"
	"strings"
)

/*
77. 组合
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
func combine(n int, k int) [][]int {
	l := make([]int, k)
	i := 0
	var ret [][]int
	for i >= 0 {
		l[i]++
		if l[i] > n-k+i+1 {
			i--
		} else if i == k-1 {
			t := make([]int, k)
			copy(t, l)
			ret = append(ret, t)
		} else {
			i++
			l[i] = l[i-1]
		}
	}
	return ret
}

func combine(n int, k int) [][]int {

	var arr [][]int
	var hashMap = make(map[string]struct{})
	Combine(1, n, k, hashMap, &arr, "")

	return arr
}

func Combine(cur, n, k int, hashMap map[string]struct{}, result *[][]int, curStr string) {
	if k == 0 {
		if _, ok := hashMap[curStr]; !ok {
			hashMap[curStr] = struct{}{}
			var curArr []int
			for _, v := range strings.Split(curStr, " ") {
				if v == "" {
					continue
				}
				n, _ := strconv.Atoi(v)
				curArr = append(curArr, n)
			}
			*result = append(*result, curArr)
		}
		return
	}
	if k < 0 {
		return
	}

	for i := cur; i <= n; i++ {
		Combine(i+1, n, k-1, hashMap, result, curStr+" "+strconv.Itoa(i))
	}
}
