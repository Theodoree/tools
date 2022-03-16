package _00_699

/*
605. 种花问题

假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数 n 。能否在不打破种植规则的情况下种入 n 朵花？能则返回True，不能则返回False。

示例 1:

输入: flowerbed = [1,0,0,0,1], n = 1
输出: True
示例 2:

输入: flowerbed = [1,0,0,0,1], n = 2
输出: False
*/

func canPlaceFlowers(flowerbed []int, n int) bool {

    switch  {
    case len(flowerbed) >= 2:
        if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
            flowerbed[len(flowerbed)-1] = 1
            n--
        }

        if flowerbed[0] == 0 && flowerbed[1] == 0 {
            flowerbed[0] = 1
            n--
        }
    case len(flowerbed) == 1 && n == 1 && flowerbed[0] == 0:
        return  true
    }



    for i := 1; i < len(flowerbed)-2; i++ {
        if n == 0 {
            return true
        }

        if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
            flowerbed[i] = 1
            n--
        }
    }
    return n <= 0
}
