package _00_299

/*
243. 最短单词距离

给定一个单词列表和两个单词 word1 和 word2，返回列表中这两个单词之间的最短距离。

示例:
假设 words = ["practice", "makes", "perfect", "coding", "makes"]

输入: word1 = “coding”, word2 = “practice”
输出: 3
输入: word1 = "makes", word2 = "coding"
输出: 1
注意:
你可以假设 word1 不等于 word2, 并且 word1 和 word2 都在列表里。




public int shortestDistance(String[] words, String word1, String word2) {
    int i1 = -1, i2 = -1;
    int minDistance = words.length;
    int currentDistance;
    for (int i = 0; i < words.length; i++) {
        if (words[i].equals(word1)) {
            i1 = i;
        } else if (words[i].equals(word2)) {
            i2 = i;
        }

        if (i1 != -1 && i2 != -1) {
            minDistance = Math.min(minDistance, Math.abs(i1 - i1));
        }
    }
    return minDistance;
}
*/

func shortestDistance(words []string, word1 string, word2 string) int {
    i1 := -1
    i2 := -1

    minDistance := len(words)

    for i := 0; i < len(words); i++ {
        if words[i] == word1{
            i1 = i
        } else if words[i] == word2 {
            i2 = i
        }

        if i1 != -1 && i2 != -1 {
            result := int(math.Abs(float64(i1 - i2)))
            if minDistance > result {
                minDistance = result
            }
        }
    }

    return minDistance
}
