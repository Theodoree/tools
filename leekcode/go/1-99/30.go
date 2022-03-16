package __99

func findSubstring(s string, words []string) []int {

	for i := 0; i < len(words); i++ {

		findSubstring(s, GetWord(words, words[i]))

	}

	findSubstring(s, GetWord)

}

func GetWord(s []string) []string {

	var words []string

	index := 0
	for index < len(words) {

		for i := 0; i < len(s); i++ {

		}

	}

}

func main() {

}
