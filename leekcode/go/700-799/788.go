package _00_799

var s = []int{
    0, 0, 1, -1, -1, 1, 1, -1, 0, 1,
}

func rotatedDigits(N int) int {

    var count int
    for i := 0; i <= N; i++ {

        if isGoods(i) {
            count++
        }

    }
    return count
}

func isGoods(i int) bool {
    flag := 0
    for i > 0 {
        mod := i % 10
        flag |= s[mod]
        i /= 10
    }

    return flag > 0

}

func main() {

    rotatedDigits(2)

}
