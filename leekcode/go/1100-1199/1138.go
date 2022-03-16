package _100_1199



/*
1138. 字母板上的路径

我们从一块字母板上的位置 (0, 0) 出发，该坐标对应的字符为 board[0][0]。

在本题里，字母板为board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"].

我们可以按下面的指令规则行动：

如果方格存在，'U' 意味着将我们的位置上移一行；
如果方格存在，'D' 意味着将我们的位置下移一行；
如果方格存在，'L' 意味着将我们的位置左移一列；
如果方格存在，'R' 意味着将我们的位置右移一列；
'!' 会把在我们当前位置 (r, c) 的字符 board[r][c] 添加到答案中。
返回指令序列，用最小的行动次数让答案和目标 target 相同。你可以返回任何达成目标的路径。



示例 1：

输入：target = "leet"
输出："DDR!UURRR!!DDD!"
示例 2：

输入：target = "code"
输出："RR!DDRR!UUL!R!"
*/

func alphabetBoardPath(target string) string {

    var result string
    var last int32
    last = 'a'
    for _, v := range target {
        for {
            if (v-'a')/5 > (last-'a')/5 && last+5 <= 'z' {
                last += 5
                result += "D"
            } else if (v-'a')/5 < (last-'a')/5 {
                last -= 5
                result += "U"
            } else {
                break
            }
        }
        var k bool
        for  {

            if v == 'z' && v != last {
                v = 'u'
                k = true
            }

            if v > last {
                last++
                result += "R"
            } else if v < last {
                last--
                result += "L"
            } else {
                if k {
                    result+="D"
                    last = 'z'
                    k = false
                }
                break
            }
        }
        result += "!"
    }

    return result
}