package _00_799

/*
734. 句子相似性
给定两个句子 words1, words2 （每个用字符串数组表示），和一个相似单词对的列表 pairs ，判断是否两个句子是相似的。
*/
func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}

	var m = make(map[string]map[string]struct{})

	for _, v := range similarPairs {
		if _m, ok := m[v[0]]; !ok {
			m[v[0]] = map[string]struct{}{v[1]: {}}
		} else {
			_m[v[1]] = struct{}{}
		}

		if _m, ok := m[v[1]]; !ok {
			m[v[1]] = map[string]struct{}{v[0]: {}}
		} else {
			_m[v[0]] = struct{}{}
		}
	}

	for i := 0; i < len(sentence1); i++ {
		if sentence1[i] == sentence2[i] {
			continue
		}

		if _m, ok := m[sentence1[i]]; !ok {
			return false
		} else if _, ok = _m[sentence2[i]]; !ok {
			return false
		}
	}

	return true

}
