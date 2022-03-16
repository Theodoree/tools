package _700_1799


/*
1772. 按受欢迎程度排列功能
给定一个字符串数组 features ，其中 features[i] 是一个单词，描述你最近参与开发的项目中一个功能的名称。你调查了用户喜欢哪些功能。
另给定一个字符串数组 responses，其中 responses[i] 是一个包含以空格分隔的一系列单词的字符串。
你想要按照受欢迎程度排列这些功能。
严格地说，令 appearances(word) 是满足 responses[i] 中包含单词 word 的 i 的个数，则当 appearances(features[x]) > appearances(features[y]) 时，第 x 个功能比第 y 个功能更受欢迎。
返回一个数组 sortedFeatures ，包含按受欢迎程度排列的功能名称。当第 x  个功能和第 y 个功能的受欢迎程度相同且 x < y 时，你应当将第 x 个功能放在第 y 个功能之前。
*/
func sortFeatures(features []string, responses []string) []string {
	var m  = make(map[string]map[int]struct{})
	for _,v:=range features{
		m[v] = make(map[int]struct{})
	}

	for idx:=range responses{
		arr:=strings.Split(responses[idx]," ")
		for _,v:=range arr{
			if sub,ok:=m[v];ok{
				sub[idx] = struct{}{}
			}
		}

	}

	sort.SliceStable(features, func(i, j int) bool {
		return len(m[features[i]]) > len(m[features[j]])
	})

	return features
}



