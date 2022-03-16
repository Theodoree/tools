package _00_299


/*
293. 翻转游戏

你和朋友玩一个叫做「翻转游戏」的游戏，游戏规则：给定一个只有 + 和 - 的字符串。你和朋友轮流将 连续 的两个 "++" 反转成 "--"。 当一方无法进行有效的翻转时便意味着游戏结束，则另一方获胜。

请你写出一个函数，来计算出每个有效操作后，字符串所有的可能状态。

示例：

输入: s = "++++"
输出:
[
  "--++",
  "+--+",
  "++--"
]
注意：如果不存在可能的有效操作，请返回一个空列表 []。
*/

func generatePossibleNextMoves(s string) []string {

    var result []string
    f := []byte(s)
    for i := 0; i < len(s)-1; i++ {
        f = []byte(s)
        if f[i] == f[i+1] && f[i] == '+' {
            f[i] = '-'
            f[i+1] = '-'
            result = append(result, string(f))
        }

    }

    return result

}