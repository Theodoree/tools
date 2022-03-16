package _00_499

import "fmt"

/*
414. 第三大的数
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
*/
func thirdMax(nums []int) int {

    var first, second, third = -1, -1, -1

    for _, v := range nums {

        if v > third || third == -1 {
            if v != second && v != first {

                if third != -1 && second == -1 {
                    if third > v {
                        second = third
                        third = v
                    } else {
                        second = v
                    }
                } else {
                    third = v
                }

            }
        }

        if second < third {
            third, second = second, third
        }

        if first < second {
            first, second = second, first
        }

    }

    if third == -1 || second == -1 {
        return first
    }
    return third
}
/*
func thirdMax1(nums []int) int {
    n := len(nums)
    if n <= 0 {
        return 0
    }

    var max = [3]int{-0xffffffff, -0xffffffff, -0xffffffff}
    var maxCount [3]int

    max[0],  maxCount[0] = nums[0], 1

    for i:=1; i<n; i++ {
        x := nums[i]

        if x > max[0] {
            max[2], maxCount[2] = max[1], maxCount[1]
            max[1], maxCount[1] = max[0], maxCount[0]
            max[0], maxCount[0] = x, 1
        } else if x == max[0] {
            maxCount[0]++
        } else if x > max[1] {
            max[2], maxCount[2] = max[1], maxCount[1]
            max[1], maxCount[1] = x, 1
        } else if  x == max[1] {
            maxCount[1]++
        } else if  x > max[2] {
            max[2], maxCount[2] = x, 1
        } else if  x == max[2] {
            maxCount[2]++
        }
    }

    if maxCount[2] > 0 {
        return max[2]
    }
    return max[0]
}

func max(x,y int) int {
    if x > y {
        return x
    }
    return y
}
*/

func main() {
    fmt.Println(thirdMax([]int{-2147483648,1,1}))

}
