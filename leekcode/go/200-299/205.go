package _00_299

import (
    "strconv"
    "strings"
)

/*
205. 同构字符串

给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。

所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

示例 1:

输入: s = "egg", t = "add"
输出: true
示例 2:

输入: s = "foo", t = "bar"
输出: false
示例 3:

输入: s = "paper", t = "title"
输出: true
*/

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    filter := make(map[string]int)

    s1 := strings.Split(s, ``)
    t1 := strings.Split(t, "")

    for k, v := range s1 {
        val, ok := filter[v]
        if ok {
            s1[k] = strconv.Itoa(val)
        } else {
            filter[v] = len(filter)
            s1[k] = strconv.Itoa(len(filter))
        }
    }

    filter = make(map[string]int)

    for k, v := range t1 {
        val, ok := filter[v]
        if ok {
            t1[k] = strconv.Itoa(val)
        } else {
            filter[v] = len(filter)
            t1[k] = strconv.Itoa(len(filter))
        }
    }
    s = strings.Join(s1,``)
    t = strings.Join(t1,``)
    return s == t
}

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var m = map[rune]byte{}
    var n = map[byte]int{}
    for i, char := range s {
        if v, ok := m[char]; !ok {
            m[char] = t[i]
            if _, ok := n[t[i]]; !ok {
                n[t[i]]++
            } else {
                return false
            }
        } else if v != t[i] {
            return false
        }
    }
    return true
}