package _100_1199

/*
1128. 等价多米诺骨牌对的数量
给你一个由一些多米诺骨牌组成的列表 dominoes。
如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。
形式上，dominoes[i] = [a, b] 和 dominoes[j] = [c, d] 等价的前提是 a==c 且 b==d，或是 a==d 且 b==c。
在 0 <= i < j < dominoes.length 的前提下，找出满足 dominoes[i] 和 dominoes[j] 等价的骨牌对 (i, j) 的数量。
*/
func numEquivDominoPairs(dominoes [][]int) int {
	var buf [100]int
	var ans int
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i][0] > dominoes[i][1] {
			dominoes[i][0], dominoes[i][1] = dominoes[i][1], dominoes[i][0]
		}

		ans += buf[dominoes[i][0]*10+dominoes[i][1]]
		buf[dominoes[i][0]*10+dominoes[i][1]]++
	}
	return ans
}
