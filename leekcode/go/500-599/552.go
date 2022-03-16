package _00_599

/* 回溯算法

var t = map[int]string{
    1: `A`,
    2: `L`,
    3: `P`,
}

func checkRecord(n int) int {
    var sum int
    CheckRecord(``, 1, n, &sum)
    return sum
}

func CheckRecord(current string, cnt, limit int, n *int) {
    if len(current) >= limit {
        if Check10(current) {
            *n++
        }
        return
    }

    for _, v := range t {
        CheckRecord(current+v, cnt+1, limit, n)
    }

}

//不能多于一个'A'
//不能超过两个连续的'L'（迟到）

func Check10(s string) bool {

    var Acnt int
    var lCnt int

    for i := 0; i < len(s); i++ {
        if s[i] == 'L' {
            lCnt++
        } else {
            lCnt = 0
        }

        if s[i] == 'A' {
            Acnt++
        }
        if Acnt > 1 || lCnt > 2{
            return false
        }
    }
    return true
}

先考慮只有兩個字母PL的情況，不出現連續三個L的情況個數是f(n)的話，那麼有通項公式f(n)=f(n-1)+f(n-2)+f(n-3)，
然後A最多只有一個，可以插入n-1個PL鐘任意位置，把之前的序列分為兩個序列，然後可能的情況是f(i)*f(n-i-1)

*/

//动态规划算法
//L A P       A只能使用一次 L不能连续
//3 8 19 43 94 200 418 861

func main() {

}
