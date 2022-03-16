package _go

import "math"

/*
面试题 05.04. 下一个数
下一个数。给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）。
示例1:
 输入：num = 2（或者0b10）
 输出：[4, 1] 或者（[0b100, 0b1]）
示例2:
 输入：num = 1
 输出：[2, -1]
*/

func findClosedNumbers(num int) []int {

	switch num {
	case 1:
		return []int{2, -1}
	case math.MaxInt32:
		return []int{-1, -1}
	}

	return []int{getLarge(num), getMin(num)}
}

func getLarge(num int) int {
	var firstInt int
	firstInt = -1
	var j int

	for i := 0; i < 64; i++ {
		if num&(1<<i) == 0 { // 空cell
			if firstInt == -1 {
				continue
			}

			num += 1 << i
			num -= 1 << firstInt
			j = i
			break
		} else if firstInt == -1 {
			firstInt = i
		}
	}

	var cnt int
	for i := 0; i < j; i++ {
		if num&(1<<i) > 0 {
			num -= 1 << i
			cnt++
		}
	}

	for i := 0; i < 64; i++ {
		if cnt == 0 {
			break
		}
		if num&(1<<i) == 0 {
			num += 1 << i
			cnt--
		}
	}

	return num
}

func getMin(num int) int {
	var firstInt int
	firstInt = -1
	var j int
	var cnt int
	for i := 0; i < 64; i++ {
		if num&(1<<i) > 0 {
			if firstInt == -1 {
				num -= 1 << i
				cnt++
				continue
			}
			num -= 1 << i
			num += 1 << firstInt
			j = i
			break
		} else {
			firstInt = i
		}
	}

	for i := j - 1; i >= 0; i-- {
		if cnt == 0 {
			break
		}
		if num&(1<<i) == 0 {
			num += 1 << i
			cnt--
		}
	}

	return num
}
