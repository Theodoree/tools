package _go

import "sort"

/*
剑指 Offer 61. 扑克牌中的顺子
从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。



示例 1:

输入: [1,2,3,4,5]
输出: True


示例 2:

输入: [0,0,1,2,5]
输出: True
*/

func isStraight(nums []int) bool {

	sort.Ints(nums)

	var joker int

	var lastNumber int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			joker++
			continue
		}

		if lastNumber > 0 && nums[i] != lastNumber+1 {
			for lastNumber != nums[i]-1 {
				if joker > 0 {
					lastNumber++
					joker--
				} else {
					return false
				}
			}
		}

		lastNumber = nums[i]

	}

	return true

}
