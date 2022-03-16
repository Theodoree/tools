package _00_799

/*
744. 寻找比目标字母大的最小字母

给定一个只包含小写字母的有序数组letters 和一个目标字母 target，寻找有序数组里面比目标字母大的最小字母。

数组里字母的顺序是循环的。举个例子，如果目标字母target = 'z' 并且有序数组为 letters = ['a', 'b']，则答案返回 'a'。

示例:

输入:
letters = ["c", "f", "j"]
target = "a"
输出: "c"

输入:
letters = ["c", "f", "j"]
target = "c"
输出: "f"

输入:
letters = ["c", "f", "j"]
target = "d"
输出: "f"

输入:
letters = ["c", "f", "j"]
target = "g"
输出: "j"

输入:
letters = ["c", "f", "j"]
target = "j"
输出: "c"

输入:
letters = ["c", "f", "j"]
target = "k"
输出: "c"
*/
func nextGreatestLetter(letters []byte, target byte) byte {

    l, r := 0, len(letters)-1
    var min byte
    for l < r {

        mid := l + (r-l)/2

        if r == mid || l == mid {
            break
        }

        if letters[mid] > target {
            r = mid
        } else if letters[mid] < target {
            l = mid
        } else {
            break
        }

    }
    for i := l; i < len(letters); i++ {
        if letters[i] > target {
            min = letters[i]
            break
        }

    }

    if min == 0 {
        min = letters[0]
    }

    return min
}
