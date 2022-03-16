package _00_399

import "math/rand"

/*
398. 随机数索引
给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。
注意：
数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。
*/

type Solution struct {
    nums []int
}


func Constructor(nums []int) Solution {
    return Solution{nums: nums}
}


func (this *Solution) Pick(target int) int {
    var c int64 =0
    idx:=0
    for i:=0;i<len(this.nums);i++{
        if this.nums[i] == target{
            c++
            if rand.Int63()%c == 0 {
                idx = i
            }
        }
    }
    return idx

}
