package _00_299

/*
273. 整数转换英文表示。
*/

var single = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
var teens = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var double = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var others = []string{"Hundred", "Thousand", "Million", "Billion", "Trillion"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	arr := numToSlice(num)
	var result []string
	for i := 0; i < len(arr); i++ {
		str := NumberToWords(arr[i])
		result = append(result, str...)
		if len(str) > 0 {
			switch len(arr) - i - 1 {
			case 1:
				result = append(result, others[1])
			case 2:
				result = append(result, others[2])
			case 3:
				result = append(result, others[3])
			}
		}
	}

	return strings.Join(result, " ")
}

func numToSlice(num int) []int {
	var arr []int
	for num > 0 {
		arr = append(arr, num%1000)
		num /= 1000
	}

	if len(arr) == 1 {
		return arr
	}

	var r = make([]int, 0, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		r = append(r, arr[i])
	}

	return r
}

func NumberToWords(num int) []string {
	if num == 0 {
		return nil
	}

	var result []string

	var nums [3]int
	nums[0] = (num % 1000) / 100
	nums[1] = (num % 100) / 10
	nums[2] = num % 10

	if nums[0] > 0 {
		result = append(result, single[nums[0]])
		result = append(result, others[0])
	}

	switch nums[1] {
	case 1:
		result = append(result, teens[nums[2]])
	default:
		if nums[1] > 0 {
			result = append(result, double[nums[1]])
		}
		if nums[2] > 0 {
			result = append(result, single[nums[2]])
		}
	}

	return result
}
