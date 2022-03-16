package _00_899

func buddyStrings(A string, B string) bool {

    /*
       字符串长度不相等, 直接返回false
       字符串相等的时候, 只要有重复的元素就返回true
       A, B字符串有不相等的两个地方, 需要查看它们交换后是否相等即可.
    */
    if len(A) != len(B) {
        return false
    }

    if A == B {

        c := [26]int{}
        for _, a := range A {
            c[a-'a']++
            if c[a-'a'] >= 2 {
                return true
            }
        }
        return false
    }

    i := 0
    countDown := 2
    ca, cb := byte(0), byte(0)
    for countDown > 0 && i < len(A) {
        if A[i] != B[i] {
            // 这里等于将不同的字母都相加,最多有两个不同
            ca += A[i]
            cb += B[i]
            countDown--
        }

        i++
    }

    // 这里判断 ca == cb,同时判断往后的字符串必须相等
    return ca == cb && A[i:] == B[i:]

}
