package __99

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {

	if dividend > math.MaxInt32 || dividend < math.MinInt32 {
		return 0
	}

	if divisor > math.MaxInt32 || divisor < math.MinInt32 {
		return 0
	}

	if divisor == 0 || dividend == 0 {
		return 0
	}
	//通过异或来判断是否符号相异
	negative := (dividend ^ divisor) < 0
	var result int
	t := math.Abs(float64(dividend))
	d := math.Abs(float64(divisor))
	for i := 31; i >= 0; i-- {

		if int(t)>>uint8(i) >= int(d) {
			result += 1 << uint8(i)
			t = t - float64(int(d)<<uint8(i))
		}

	}

	if negative {
		result = -result
		if result < math.MinInt32{
			result = math.MinInt32
		}
	}

	if result > math.MaxInt32{
		result = math.MaxInt32
	}

	return result
}

func main() {

	fmt.Println(divide(10, 3)) //3
	fmt.Println(divide(7, -3)) //-2
	fmt.Println(1^1)

}
