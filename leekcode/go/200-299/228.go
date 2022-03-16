package _00_299

import "fmt"

func summaryRanges(nums []int) []string {

    var l, r int
    var result []string

    for len(nums) > r {

        for r+1 < len(nums) && nums[r+1 ]-nums[r] == 1 {
            r++
        }

        if l < r {
            result = append(result, fmt.Sprintf("%d->%d", nums[l], nums[r]))
        } else {
            result = append(result, fmt.Sprintf("%d", nums[l]))
        }

        r++
        l = r
    }

    return result
}

func main() {

    summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})

}
