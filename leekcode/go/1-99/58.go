package __99

import (
	"fmt"
	"strings"
)

// 58. 最后一个单词的长度

func lengthOfLastWord(s string) int {

	f := strings.Split(s, ` `)

	var length int

	for i:=0;i<len(f);i++{

		if len(f[i]) > 0 {
			length = len(f[i])
		}

	}

	return length
}

func main() {

	fmt.Println(lengthOfLastWord(`a `))


}
