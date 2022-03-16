package __99

import "fmt"

/*
9.回文数
*/

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	var sum int
	for tmp := x; tmp > 0; tmp /= 10 {
		tail := tmp % 10
		sum = sum*10 + tail
	}
	return x == sum
}

func main() {
	fmt.Println(isPalindrome(121))

}
