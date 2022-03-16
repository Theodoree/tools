package _00_299



func getHint(secret string, guess string) string {

    secretM := make(map[byte]int)
    var b, c int

    var gByte, sByte []byte

    for idx, v := range []byte(secret) {
        if v == guess[idx] {
            b++
        } else {
            gByte = append(gByte, guess[idx])
            sByte = append(sByte, secret[idx])
            secretM[v]++
        }
    }

    for i := 0; i < len(gByte); i++ {
        if val, ok := secretM[gByte[i]]; ok {
            c++
            if val == 1 {
                delete(secretM, gByte[i])
            } else {
                secretM[gByte[i]]--
            }
        }
    }

    return fmt.Sprintf("%dA%dB", b, c)
}
