package __99

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	var sum int
	var flag bool
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}

		switch {
		case str[i] == '-':
			flag = true
			i++
		case str[i] == '+':
			i++
		case str[i] >= '0' && str[i] <= '9':
		default:
			return 0
		}

		for j := i; j < len(str); j++ {
			if str[j] >= '0' && str[j] <= '9' {
				if !flag {
					sum = sum*10 + int(str[j]-'0')
					if sum > math.MaxInt32 {
						return math.MaxInt32
					}
				} else {
					sum = sum*10 - int(str[j]-'0')
					if sum < math.MinInt32 {
						return math.MinInt32
					}
				}
				continue
			}
			break
		}

		return sum

	}

	return sum

}

func main() {

	fmt.Println(myAtoi(`+1`))

}
