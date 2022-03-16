package _00_799

import (
	"fmt"
)

//  771. 宝石与石头

func numJewelsInStones(J string, S string) int {
	ret := 0
	for _, c := range J {
		for _, i := range S {
			if c == i {
				ret++
			}
		}
	}
	return ret
}

func main() {

	fmt.Println(numJewelsInStones(`aA`, `aAAbbbb`))
}
