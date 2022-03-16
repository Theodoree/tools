package _00_699



/*
628. 三个数的最大乘积

给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1:

输入: [1,2,3]
输出: 6
示例 2:

输入: [1,2,3,4]
输出: 24
注意:

给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
*/
func maximumProduct(nums []int) int {
    sort.Ints(nums)
    i :=math.Max(float64(nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3]),float64(nums[0]*nums[1]*nums[len(nums)-1]) )
    return  int(i)
}

func maximumProduct(nums []int) int {
    first := -1000
    second := -1000
    third := -1000
    fourth := 1000
    fifth := 1000
    for i:=0;i<len(nums);i++ {
        if nums[i] > first {
            third = second
            second = first
            first = nums[i]
        } else if nums[i] > second {
            third = second
            second = nums[i]
        } else if nums[i] > third {
            third = nums[i]
        }

        if nums[i] < fifth {
            fourth = fifth
            fifth = nums[i]
        } else if nums[i] < fourth {
            fourth = nums[i]
        }
    }

    if first * second * third > fourth * fifth * first {
        return first * second * third
    }

    return fourth * fifth * first
}
