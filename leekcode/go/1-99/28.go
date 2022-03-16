package __99

import (
	"fmt"
)

func strStr(haystack string, needle string) int {

	n := len(needle)

	switch {
	case n == 0:
		return 0
	case n == len(haystack):
		if haystack == needle {
			return 0
		}
		return -1
	case len(haystack) < n:
		return -1

	}

	var flag bool
	for i := 0; i < len(haystack); i++ {
		if len(haystack)-i >= n && haystack[i] == needle[0] {
			str := haystack[i:]
			for i := 0; i < len(needle) && len(str) >= len(needle); i++ {
				if str[i] != needle[i] {
					break
				}
				if i == len(needle)-1 {
					flag = true
				}
			}
			if flag {
				return i
			}
		}
	}

	return -1

}

func main() {
	f := strStr(`mississippi`, `pi`)
	fmt.Println(f)

}
