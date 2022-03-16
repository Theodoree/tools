package _00_199

//190. 颠倒二进制位
func reverseBits(num uint32) uint32 {
    var sum int
    s := int(num)
    for i := 0; i < 32; i++ {
        sum += ((s >> uint(i)) & 1) << (31 - uint(i))
    }
    return uint32(sum)
}

