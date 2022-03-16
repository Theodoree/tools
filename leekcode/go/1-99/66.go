package __99

import "fmt"

//66.加一

func plusOne(digits []int) []int {
	var carry int = 1
	var l = len(digits)
	for i := l-1; i>=0; i-- {
		digits[i] = digits[i] + carry
		if digits[i] == 10 {
			digits[i] = 0
			carry =1
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {

	f := plusOne([]int{9,9})
	fmt.Println(f)

}
