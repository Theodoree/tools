package __99

import (
	"fmt"
)

var hash = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	lens := len(digits)
	if lens == 0 {
		return nil
	}
	var res []string
	if lens >= 1 {
		s := digits[0]

		recur := letterCombinations(digits[1:])
		fmt.Println("recur ",recur)

		if recur != nil {
			for _, str := range recur {
				for _, ch := range hash[s] {
					res = append(res, ch+str)
				}
			}
			return res
		} else {
			return hash[s]
		}
	}
	return res
}

func main() {
	fmt.Println(letterCombinations(`2425`))
}
