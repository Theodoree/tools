package __99

import "fmt"

var Int_Roma = map[int]string{
	1:    `I`,
	5:    `V`,
	10:   `X`,
	50:   `L`,
	100:  `C`,
	500:  `D`,
	1000: `M`,
}

func intToRoman(num int) string {

	var val string
	slice := []int{1000, 100, 10, 1}
	T := num / 1000
	num %= 1000
	slice = slice[1:]
	for i := 0; i < T; i++ {
		val += Int_Roma[1000]
	}

	for len(slice) > 0 {
		key := slice[0]

		t := num / key
		num %= key

		if t == 9 {
			val += Int_Roma[key] + Int_Roma[key*10]
		} else if t == 4 {
			val += Int_Roma[key] + Int_Roma[key*5]
		} else {
			if t >= 5 {
				t -= 5
				val += Int_Roma[key*5]
			}

			for i := 0; i < t; i++ {
				val += Int_Roma[key]
			}
		}

		slice = slice[1:]

	}

	return val
}

func main() {

	fmt.Println(intToRoman(1994))

}
