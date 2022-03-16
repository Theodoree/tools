package __99

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {

	var cnt string
	var count int

	if len(a) > len(b) {
		a, b = b, a
	}
	count = len(b) - len(a)
	//补位
	for i := 0; i < count; i++ {
		a = `0` + a
	}


	var carry int

	for i := len(a) - 1; i >= 0; i-- {
		val1,_ :=strconv.Atoi(string(a[i]))
		val2,_ :=strconv.Atoi(string(b[i]))
		sum := val1 + val2 + carry
		switch {
		case sum == 3:
			cnt = `1` + cnt
			carry = 1
		case sum == 2:
			cnt = `0` + cnt
			carry = 1
		case sum == 1:
			cnt = `1` + cnt
			carry = 0
		case sum == 0:
			cnt = `0` + cnt
			carry = 0

		}

	}

	if carry > 0 {
		cnt = `1` + cnt
	}

	return cnt

}

func main() {

	f := addBinary(`1010`, `1011`)
	fmt.Println(f)
}
