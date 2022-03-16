package _00_799


//747.至少是其他数字两倍的最大数
func dominantIndex(nums []int) int {

    switch {
    case len(nums) == 0:
        return -1
    case len(nums) == 1:
        return 0
    }

    s := make([]int, 2)
    s[0] = -1
    s[1] = -1
    var max, second int

    for i := 0; i < len(nums); i++ {
        if nums[i] > s[1] {
            s[1] = nums[i]
            second = i
        }

        if s[1] > s[0] {
            s[0], s[1] = s[1], s[0]
            max,second= second,max
        }
    }
    if nums[max] >= nums[second]*2 {
        return max
    }

    return -1

}

/*
func dominantIndex(nums []int) int {
    firMax := -1
    secMax := -1
    index := -1
    for i ,v  := range nums {
        if v > firMax {
            secMax = firMax
            firMax = v
            index  = i

        } else if v > secMax {
            secMax = v
        }
    }

    if firMax >= secMax*2 {
        return index
    }
    return -1
}
*/
