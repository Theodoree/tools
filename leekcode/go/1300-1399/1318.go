package _300_1399



/*
1318. 或运算的最小翻转次数
给你三个正整数 a、b 和 c。
你可以对 a 和 b 的二进制表示进行位翻转操作，返回能够使按位或运算   a OR b == c  成立的最小翻转次数。
「位翻转操作」是指将一个数的二进制表示任何单个位上的 1 变成 0 或者 0 变成 1 。
*/
func minFlips(a int, b int, c int) int {
	if a | b == c {
		return 0
	}

	var  cnt int
	for i:=0;i<32;i++{
		if 1<<i > a && 1 <<i > b {
			break
		}

		if (1<<i)&a > 0   && (1<<i)&b > 0  && (1<<i)&c == 0 {
			cnt++
		}
	}

	sum:=a|b
	for i:=0;i<32;i++{
		if (1<<i)&c  != (1<<i)&sum {
			cnt++
		}
	}

	return cnt



}
