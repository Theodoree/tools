package _100_1199


/*
1119. 删去字符串中的元音

给你一个字符串 S，请你删去其中的所有元音字母（ 'a'，'e'，'i'，'o'，'u'），并返回这个新字符串。

示例 1：

输入："leetcodeisacommunityforcoders"
输出："ltcdscmmntyfrcdrs"
示例 2：

输入："aeiou"
输出：""


提示：

S 仅由小写英文字母组成。
1 <= S.length <= 1000
*/

var FilterMap = map[byte]struct{}{
    'a': struct{}{},
    'e': struct{}{},
    'i': struct{}{},
    'o': struct{}{},
    'u': struct{}{},
}

func removeVowels(S string) string {

    var result string
    for _, v := range S {
        if _, ok := FilterMap[byte(v)]; ok {
            continue
        }

        result += string(v)

    }
    return result
}