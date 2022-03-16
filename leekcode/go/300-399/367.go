package _00_399



/*
367. 有效的完全平方数


给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。

说明：不要使用任何内置的库函数，如  sqrt。

示例 1：

输入：16
输出：True
示例 2：

输入：14
输出：False
*/

//二分
func isPerfectSquare(num int) bool {
    l, r := 1, num
    for l <= r {
        m := l + (r-l)/2

        if m*m < num {
            l = m + 1
        } else if m*m > num {
            r = m - 1
        } else {
            return true
        }
    }
    return false
}


/*
渐近测试
func isPerfectSquare(num int) bool {
    i := 1
    var r int
    for i <= num {

        if r == num {
            return true
        } else {
            i++
            r = i * i
        }
    }

    return false
}


公式法
func isPerfectSquare(num int)bool {
    i := 1
    for num > 0 {
        num -= i;
        i += 2;
    }
    return num == 0;
}

*/
