package __99

import "fmt"

// 69. x 的平方根

func mySqrt(x int) int {

	if x == 0 || x == 1 {
		return x
	}

	i := (x) / 2.0

	for i*i > x {
		i = (i + x/i) / 2
	}
	return i
}

func main() {
	fmt.Println(mySqrt(8))

}
