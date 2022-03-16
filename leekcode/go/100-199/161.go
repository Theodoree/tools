package _00_199


//异或 相同为零 不同为1
// 与或& 相同为1 不同为零
func hammingDistance(x int, y int) int {

    dist := 0
    val := x ^ y
    for val > 0 {
        dist++
        val &= val - 1
    }
    return dist
}
