package _00_299

/*
231. 2的幂

给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
*/

func isPowerOfTwo(n int) bool {

    if n == 1 || n == 2 {
        return true
    }
    if n == 0 {
        return false
    }

    if n%2 != 0 {
        return false
    }

    return n&(n-1) == 0
}

func isPowerOfTwo1(n int) bool {
    if n == 1 || n == 2 {
        return true
    }

    for n > 2 {
        /*  这里判断n是否为奇数,如果为奇数肯定不是2的幂*/
        if n&1 == 1 {
            return false
        }
        /*右移计算*/
        n = n >> 1
        /*如果最终的右移等于 二进制（10） 那么就是2的幂 */
        if n == 2 {
            return true
        }
    }
    return false
}
