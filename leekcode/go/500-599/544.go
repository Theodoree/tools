package _00_599




/*
544. 输出比赛匹配对
在 NBA 季后赛中，我们总是安排较强的队伍对战较弱的队伍，例如用排名第 1 的队伍和第 n 的队伍对决，这是一个可以让比赛更加有趣的好策略。现在，给你 n 支队伍，你需要以字符串格式输出它们的 最终 比赛配对。
n 支队伍按从 1 到 n 的正整数格式给出，分别代表它们的初始排名（排名 1 最强，排名 n 最弱）。我们用括号（'(', ')'）和逗号（','）来表示匹配对——括号（'(', ')'）表示匹配，逗号（','）来用于分割。 在每一轮的匹配过程中，你都需要遵循将强队与弱队配对的原则。
*/
func findContestMatch(n int) string {
	arr:=make([]string,n)
	for i:=0;i<len(arr);i++{
		arr[i] = strconv.Itoa(i+1)
	}
	for len(arr) > 1 {
		var tmp []string
		for i:=0;i<len(arr)/2;i++{
			tmp = append(tmp,"("+arr[i]+","+arr[len(arr)-1-i]+")")
		}
		arr = tmp
	}

	return arr[0]
}

