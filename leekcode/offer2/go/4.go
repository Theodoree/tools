package _go

import "fmt"

/*
剑指 Offer II 004. 只出现一次的数字
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
*/

func singleNumber(nums []int) int {
	var a, b int
	fmt.Printf("第%d次循环 \n", 1)
	fmt.Println("a = 0 b = 0")
	for i, v := range nums {
		fmt.Printf("b = (b ^ v) & ^a = %d ^ %d & ^%d = %d & %d  \n", b, v, a, b^v, ^a)
		b = (b ^ v) & ^a
		fmt.Printf("a = (a ^ v) & ^b = %d ^ %d & ^%d = %d & %d  \n", a, v, b, a^v, ^b)
		a = (a ^ v) & ^b
		fmt.Println()
		fmt.Printf("第%d次循环 \n", i+2)
		fmt.Printf("a = %d b = %d \n", a, b)

	}
	/* 等于这里是模拟三进制,return b 的意思就是执行完第二次之后就返回
	      入参 []int{1, 1, 1, 3, 4, 4, 4, 7, 7, 7}
	   第1次循环
	   a = 0 b = 0
	   b = (b ^ v) & ^a = 0 ^ 1 & ^0 = 1 & -1
	   a = (a ^ v) & ^b = 0 ^ 1 & ^1 = 1 & -2

	   第2次循环
	   a = 0 b = 1
	   b = (b ^ v) & ^a = 1 ^ 1 & ^0 = 0 & -1
	   a = (a ^ v) & ^b = 0 ^ 1 & ^0 = 1 & -1

	   第3次循环
	   a = 1 b = 0
	   b = (b ^ v) & ^a = 0 ^ 1 & ^1 = 1 & -2
	   a = (a ^ v) & ^b = 1 ^ 1 & ^0 = 0 & -1

	   第4次循环
	   a = 0 b = 0
	   b = (b ^ v) & ^a = 0 ^ 3 & ^0 = 3 & -1
	   a = (a ^ v) & ^b = 0 ^ 3 & ^3 = 3 & -4

	   第5次循环
	   a = 0 b = 3
	   b = (b ^ v) & ^a = 3 ^ 4 & ^0 = 7 & -1
	   a = (a ^ v) & ^b = 0 ^ 4 & ^7 = 4 & -8

	   第6次循环
	   a = 0 b = 7
	   b = (b ^ v) & ^a = 7 ^ 4 & ^0 = 3 & -1
	   a = (a ^ v) & ^b = 0 ^ 4 & ^3 = 4 & -4

	   第7次循环
	   a = 4 b = 3
	   b = (b ^ v) & ^a = 3 ^ 4 & ^4 = 7 & -5
	   a = (a ^ v) & ^b = 4 ^ 4 & ^3 = 0 & -4

	   第8次循环
	   a = 0 b = 3
	   b = (b ^ v) & ^a = 3 ^ 7 & ^0 = 4 & -1
	   a = (a ^ v) & ^b = 0 ^ 7 & ^4 = 7 & -5

	   第9次循环
	   a = 3 b = 4
	   b = (b ^ v) & ^a = 4 ^ 7 & ^3 = 3 & -4
	   a = (a ^ v) & ^b = 3 ^ 7 & ^0 = 4 & -1

	   第10次循环
	   a = 4 b = 0
	   b = (b ^ v) & ^a = 0 ^ 7 & ^4 = 7 & -5
	   a = (a ^ v) & ^b = 4 ^ 7 & ^3 = 3 & -4

	   退出
	   a = 0 b = 3

	*/
	return b
}
