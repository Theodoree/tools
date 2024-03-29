package _go

/*
面试题 05.03. 翻转数位
给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。

示例 1：

输入: num = 1775(110111011112)
输出: 8
示例 2：

输入: num = 7(01112)
输出: 4
通过次数11,187提交次数29,101
*/
func reverseBits(num int) int {
	var zeroIndex, maxCount, count int
	zeroIndex = -1
	for i := 0; i < 32; i++ {
		result := 1<<i&num != 0
		if result {
			count++
		} else {
			if zeroIndex != -1 {
				zeroIndex = zeroIndex + 1
				count = i - zeroIndex
			}
			zeroIndex = i
			count++
		}

		if maxCount < count {
			maxCount = count
		}
	}

	return maxCount
}
