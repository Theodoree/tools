package _00_199

import "fmt"

//136. 只出现一次的数字

//相同为0，不同为1. 异或同一个数两次，原数不变。
func singleNumber(nums []int) int {



	var num int

	for i:=0;i<len(nums);i++{

		num ^=nums[i]
	}

	return  num
}
/*
  100
  001

  101
  010

  111
  001

  110
  010

  100
*/



func main(){
	fmt.Println(singleNumber([]int{4,1,2,1,2}))

}

