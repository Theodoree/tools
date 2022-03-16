package _00_199



/*
172. 阶乘后的零
给定一个整数 n，返回 n! 结果尾数中零的数量。

*/

func trailingZeroes(n int) int {

    var result int

    for n > 0 { //因式分解的逻辑
        result+=n/5
        n/=5
    }

    return  result
}

