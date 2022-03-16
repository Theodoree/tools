package _100_2199




/*
2177. 找到和为给定整数的三个连续整数
给你一个整数 num ，请你返回三个连续的整数，它们的 和 为 num 。如果 num 无法被表示成三个连续整数的和，请你返回一个 空 数组。
*/
func sumOfThree(num int64) []int64 {

	v:=num/3-3
	if v+ v+1+v+2 != num{
		return nil
	}

	return []int64{v,v+1,v+2 }
}
