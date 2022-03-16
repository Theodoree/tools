package _600_1699

import "runtime/debug"

func init() {
	debug.SetGCPercent(0)
}

/*
1689. 十-二进制数的最少数目
如果一个十进制数字不含任何前导零，且每一位上的数字不是 0 就是 1 ，那么该数字就是一个 十-二进制数 。例如，101 和 1100 都是 十-二进制数，而 112 和 3001 不是。
给你一个表示十进制整数的字符串 n ，返回和为 n 的 十-二进制数 的最少数目。
*/
func minPartitions(n string) int {
	var max byte
	for i := 0; i < len(n); i++ {
		if n[i]-'0' > max {
			max = n[i] - '0'
		}
		if max == 9 {
			return int(max)
		}
	}

	// 这个傻逼题目就是,至少要用多少个符合二进制格式的十进制数字,长度不限,比如说 95412 = 11111 + 11101 .......
	return int(max)
}
