package _00_199

/*
158. 用 Read4 读取 N 个字符 II
给你一个文件 file ，并且该文件只能通过给定的 read4 方法来读取，请实现一个方法使其能够使 read 读取 n 个字符。注意：你的 read 方法可能会被调用多次。
read4 的定义：
read4 API 从文件中读取 4 个连续的字符，然后将这些字符写入缓冲区数组 buf4 。
返回值是读取的实际字符数。
请注意，read4() 有其自己的文件指针，类似于 C 中的 FILE * fp 。
*/
var solution = func(read4 func([]byte) int) func([]byte, int) int {
	var buf = make([]byte, 512)
	var idx int
	n := 1
	for n != 0 {
		n = read4(buf[idx:])
		idx += n
	}
	buf = buf[:idx]
	var right int
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	return func(dist []byte, n int) int {
		if right == len(buf) {
			return 0
		}
		left := right
		right = min(len(buf), right+n)
		copy(dist, buf[left:right])
		return right - left
	}
}
