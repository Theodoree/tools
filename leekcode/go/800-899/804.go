package _00_899

// 804. 唯一摩尔斯密码词

var cryptMap = map[string]string{
	`a`: ".-",
	`b`: "-...",
	`c`: "-.-.",
	`d`: "-..",
	`e`: ".",
	`f`: "..-.",
	`g`: "--.",
	`h`: "....",
	`i`: "..",
	`j`: ".---",
	`k`: "-.-",
	`l`: ".-..",
	`m`: "--",
	`n`: "-.",
	`o`: "---",
	`p`: ".--.",
	`q`: "--.-",
	`r`: ".-.",
	`s`: "...",
	`t`: "-",
	`u`: "..-",
	`v`: "...-",
	`w`: ".--",
	`x`: "-..-",
	`y`: "-.--",
	`z`: "--..",
}

func uniqueMorseRepresentations(words []string) int {
	hashmap := make(map[string]struct{})
	var cnt int

	for _, word := range words {

		var str string
		for i := 0; i < len(word); i++ {
			str += cryptMap[string(word[i])]
		}
		if _, ok := hashmap[str]; ok {
			continue
		}
		hashmap[str] = struct{}{}
		cnt++
	}

	return cnt
}
