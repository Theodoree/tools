package __99

import (
    "fmt"
    "sort"
)

//三数之和
func threeSum(nums []int) [][]int {
    L := len(nums)
    if L < 3 {
        return [][]int{}
    }
    sort.Ints(nums)
    m := make(map[[3]int]bool)

    left, right := 0, L-1
    for i := 0; i < L; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        left, right = i+1, L-1
        for left < right {
            if left == i {
                left++
                continue
            }

            if left != i+1 && nums[left] == nums[left-1] {
                left++
                continue
            }

            if right == i || (right < L-1 && nums[right] == nums[right+1]) {
                right--
                continue
            }

            sum := nums[i] + nums[left] + nums[right]
            if sum < 0 {
                left++
            } else if sum > 0 {
                right--
            } else {
                var o [3]int
                if nums[i] <= nums[left] {
                    o = [3]int{nums[i], nums[left], nums[right]}
                } else if nums[i] > nums[right] {
                    o = [3]int{nums[left], nums[right], nums[i]}
                } else {
                    o = [3]int{nums[left], nums[i], nums[right]}
                }

                m[o] = true

                left++
                right--
            }

        }
    }

    var ret [][]int

    for k := range m {
        ret = append(ret, []int{k[0], k[1], k[2]})
    }

    return ret
}

func main() {
    fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))

}
