package __99

import (
    "fmt"
    "math"
)

// 整数反转

func reverse(x int) int {
    MAX := math.MaxInt32
    MIN := math.MinInt32
    sum := 0
    plus := true
    if x < 0 {
        plus = false
    }
    for x != 0 {
        tail := x % 10
        x /= 10
        sum = sum*10 + tail
        if plus && sum > MAX {
            return 0
        } else if sum < MIN {
            return 0
        }
    }
    return sum
}

func main() {
    // reverse(1234)
    
    fmt.Println(123 % 10)
}
