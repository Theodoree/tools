package __99

import "fmt"

/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
0 红
1 白
2 蓝



一个直观的解决方案是使用计数排序的两趟扫描算法。
首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
*/
func sortColors(nums []int) {

    var Red, White, Bule int
    for i := 0; i < len(nums); i++ {
        switch nums[i] {
        case 0:
            Red++
        case 1:
            White++
        case 2:
            Bule++
        }
    }

    var cnt int
    for {
        if Red > 0 {
            nums[cnt] = 0
            cnt++
            Red--
        } else if White > 0 {
            nums[cnt] = 1
            cnt++
            White--
        } else if Bule > 0 {
            nums[cnt] = 2
            cnt++
            Bule--
        } else {
            break
        }
    }

}


func main() {
    nums := []int{2, 0, 2, 1, 1, 0}
    fmt.Println(nums)

}
