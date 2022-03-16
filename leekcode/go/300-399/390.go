package _00_399

/*
390. 消除游戏

给定一个从1 到 n 排序的整数列表。
首先，从左到右，从第一个数字开始，每隔一个数字进行删除，直到列表的末尾。
第二步，在剩下的数字中，从右到左，从倒数第一个数字开始，每隔一个数字进行删除，直到列表开头。
我们不断重复这两步，从左到右和从右到左交替进行，直到只剩下一个数字。
返回长度为 n 的列表中，最后剩下的数字。

输入:
n = 9,
1 2 3 4 5 6 7 8 9
2 4 6 8
2 6
6

输出:
6

假如输入为 n，我们使用 f(n) 表示 从左到右(forward) 的最终结果，
使用 b(n)表示 从右到左(backward) 的最终结果. f(n) = 2 * b(n/2)。
n为奇数时，b(n) = 2* f(n/2)；n为偶数是，b(n) = 2 * f(n/2) - 1.
*/

func lastRemaining(n int) int {
    if n == 1 {
        return 1
    }
    /* 暴力法 */
    /*
       var a []int
       for i := 1; i <= n; i++ {
           a = append(a, i)
       }

       bl := true
       for len(a) != 1 {
           var t []int

           if bl { //从右到左
               index := 1
               for index < len(a) {
                   t = append(t, a[index])
                   index += 2
               }
               a = t
               bl = false
           } else { //从左到右
               index := len(a) - 2
               for index >= 0 {
                   t = append(t, a[index])
                   index -= 2
               }
               sort.Ints(t)
               a = t
               bl = true
           }
       }

       return a[0]
    */
    /* 数学 */
    return 2 * (n/2 + 1 - lastRemaining(n/2))

}