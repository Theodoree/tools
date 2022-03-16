package __99

import (
	"fmt"
	"strings"
)

// 罗马数字转整数
var Roma_Int = map[string]int{
	`I`: 1,
	`V`: 5,
	`X`: 10,
	`L`: 50,
	`C`: 100,
	`D`: 500,
	`M`: 1000,
}

func romanToInt(s string) int {

	strs := strings.Split(s, ``)
	var sum int
	for i := 0; i < len(strs); i++ {

		if i+1 < len(strs) {
			if Roma_Int[strs[i]] < Roma_Int[strs[i+1]] {
				sum += Roma_Int[strs[i+1]] - Roma_Int[strs[i]]
				i++
			} else {
				num := Roma_Int[strs[i]]
				sum += num
			}
		} else {
			num := Roma_Int[strs[i]]
			sum += num
		}
	}

	return sum

}

func main() {
	fmt.Println(romanToInt(`MCMXCIV`))

}
