package __99

import (
	"fmt"
	"strings"
)

//6. Z 字形变换

/*
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/

func convert1(s string, numRows int) string {
	var res string
	d := 2*numRows - 2
	if d == 0 {
		return s
	}

	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := i; j < len(s); j += d {
				res += string(s[j])
			}
		} else {
			k2 := 2*numRows - i - 2
			for k1 := i; k1 < len(s) || k2 < len(s); {
				if k1 < len(s) {
					res += string(s[k1]);
				}
				if k2 < len(s) {
					res += string(s[k2]);
				}

				k1 += d
				k2 += d
			}
		}
	}
	return res
}


func convert(s string, numRows int) string {
	if numRows <= 1 || numRows >= len(s) {
		return s
	}
	var str = make([]string, numRows)
	i, up := -1, true
	for k, v := range s {
		if i >= numRows-1 {
			up = false
		}

		if i <= 0 {
			up = true
		}

		if up { //
			i++
		} else {
			i--
		}
		fmt.Printf("i=%d key= %d str=%s \n",i,k,string(v))
		str[i] += string(v)
	}
	fmt.Println("0 ",str[0])
	fmt.Println("1 ",str[1])
	fmt.Println("2 ",str[2])
	fmt.Println(strings.Join(str, ""))
  	return strings.Join(str, "")
}

func main(){

	convert(`LCIRETOESIIGEDHNSJHDKVJC`,3)
}