package _00_899




/*
873. 最长的斐波那契子序列的长度
如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
n >= 3
对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}
给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。
（回想一下，子序列是从原序列 arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8] 是 [3, 4, 5, 6, 7, 8] 的一个子序列）
*/

func lenLongestFibSubseq(arr []int) int {
	var m = make(map[int]struct{})
	for _,v:=range arr{
		m[v] = struct{}{}
	}
	var max int
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			key:=arr[i]+arr[j]
			last:=arr[j]
			ok:=true
			cnt:=1
			for ok {
				_,ok = m[key]
				key = last + key
				last = key - last
				cnt++

			}
			if cnt > max {
				max = cnt
			}
		}
	}
	if max < 3 {
		return 0
	}
	return max
}

