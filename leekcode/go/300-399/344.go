package _00_399

func reverseString1(s []string) {
    for i := 0; i < len(s)/2; i++ {
        j := len(s) - 1 - i
        s[i], s[j] = s[j], s[i]
    }
}

func reverseString(s []string) {

    var first, last int
    last = len(s) - 1

    for first < last {
        s[first], s[last] = s[last], s[first]
        first++
        last--
    }
}

func main() {

    reverseString([]string{"h", "e", "l", "l", "o"})
}
