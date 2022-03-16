package __99


/*
31. 下一个排列




题目描述
评论 (297)
题解(45)New
提交记录
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

*/
func nextPermutation(nums []int)  {
    l := len(nums)
    if l <= 1 {
        return
    }
    index := 0
    for i := l-2; i>=0; i-- {
        if nums[i] < nums[i+1] {
            c := l-1
            for j:=i+1;j<=l-1;j++ {
                if nums[j] <= nums[i] {
                    c = j-1
                    break
                }
            }
            nums[c],nums[i] = nums[i],nums[c]
            index = i+1
            break
        }
    }
    reverse(&nums, index, l-1)
}

func reverse (nums *[]int,l,r int) {
    for l <= r {
        (*nums)[l],(*nums)[r] = (*nums)[r],(*nums)[l]
        l++
        r--
    }
}