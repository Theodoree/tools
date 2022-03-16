package __99

import "fmt"

func removeDuplicates(nums []int) int {

    switch {
    case len(nums) == 0:
        return 0
    case len(nums) == 1:
        return 1
    }

    cnt := 1
    for i := 1; i < len(nums); i++ {
        if nums[i-1] == nums[i] {
            cnt++
        } else {
            k := i - 1
            for cnt >= 3 {
                nums[k] = -101
                cnt--
                k--
            }
            cnt = 1
        }
    }

    for cnt >= 3 {
        nums[len(nums)-(cnt-2)] = -101
        cnt--
    }
    cnt = 0
    for {
        if cnt == len(nums) {
            break
        }
        if nums[cnt] == -101 {
            nums = append(nums[:cnt], nums[cnt+1:]...)
            cnt--
        }

        cnt++
    }
    return len(nums)
}

func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }

    flag := 1
    two := false
    for index := 0; index < len(nums)-1; index++ {
        fmt.Println(nums)
        if nums[index] == nums[index+1] && !two {
            nums[flag] = nums[index+1]
            two = true
            flag++
        }
        if nums[index] != nums[index+1] {
            nums[flag] = nums[index+1]
            two = false
            flag++
        }
    }
    return flag
}


func main() {
    fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))


}
