package _00_399




/*
345. 反转字符串中的元音字母
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede"
*/
func reverseVowels(s string) string {


    var left, right int
    s_byte := []byte(s)
    right = len(s) - 1

    for left <= right {

        if isTrue(s_byte[left]) && isTrue(s_byte[right]){
            s_byte[left],s_byte[right] = s_byte[right],s_byte[left]
            left++
            right--
            continue
        }

        if !isTrue(s_byte[right]) {
            right--
            continue
        }
        if !isTrue(s_byte[left])  {
            left++
            continue
        }


    }

    return string(s_byte)

}

func isTrue(b byte) bool {
    switch b {
    case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
        return true
    }
    return false
}
