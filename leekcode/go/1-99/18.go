package __99

import (
    "sort"
)

func fourSum(nums []int, target int) [][]int {
    if len(nums) < 4 {
        return [][]int{}
    }
    sort.Ints(nums)
    var ans [][]int
    le := len(nums)
    for i := 0; i <= len(nums)-4; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
            break
        }
        if nums[i]+nums[le-1]+nums[le-2]+nums[le-3] < target {
            continue
        }
        for j := i + 1; j <= len(nums)-3; j++ {
            if j-i > 1 && nums[j] == nums[j-1] {
                continue
            }
            if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
                break
            }
            if nums[i]+nums[j]+nums[le-1]+nums[le-2] < target {
                continue
            }
            l, r := j+1, le-1
            for l < r {
                sum := nums[i] + nums[j] + nums[l] + nums[r]
                if sum == target {
                    ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
                    for l+1 < r && nums[l+1] == nums[l] {
                        l++
                    }
                    for l < r-1 && nums[r] == nums[r-1] {
                        r--
                    }
                    l, r = l+1, r-1
                } else if sum < target {
                    l++
                } else {
                    r--
                }
            }
        }
    }
    return ans
}
