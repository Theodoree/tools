package __99

import (
	"fmt"
	"strconv"
)

//43. 字符串相乘


/*
43. 字符串相乘

给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var arr []int
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	for i := len(num1) - 1; i >= 0; i-- {

		for j := len(num2) - 1; j >= 0; j-- {
			n1, _ := strconv.Atoi(string(num1[i]))
			n2, _ := strconv.Atoi(string(num2[j]))
			cnt := n1 * n2
			idx := i + j - (len(num1) + len(num2) - 2)

			if idx < 0 {
				idx = idx * -1
			}
			for len(arr) <= idx {
				arr = append(arr, 0)
			}

			if cnt >= 10 {
				for len(arr) <= idx+1 {
					arr = append(arr, 0)
				}
				arr[idx+1] += cnt / 10
			}
			arr[idx] += cnt % 10
		}
	}

	for i := 0; i < len(arr); i++ {
		cnt := 1
		if arr[i] >= 10 {
			if i+cnt >= len(arr) {
				arr = append(arr, 0)
			}

			arr[i+cnt] += arr[i] / 10
			arr[i] %= 10
			for arr[i+cnt] >= 10 {
				if i+cnt+1 >= len(arr) {
					arr = append(arr, 0)
				}
				arr [i+cnt+1] += arr[i+cnt] / 10
				arr[i+cnt] %= 10
				cnt++
			}
		}
	}

	var result string
	for k := len(arr) - 1; k >= 0; k-- {
		result += strconv.Itoa(arr[k])
	}

	return result
}

func main() {
	f := multiply(`012`, `012`)
	fmt.Println(f)
}
