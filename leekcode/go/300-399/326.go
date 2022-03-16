package _00_399

//326
func isPowerOfThree(n int) bool {
    if n <= 0 {
        return  false
    }

    return  1162261467%n == 0
}

