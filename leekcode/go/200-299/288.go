package _00_299

import (
    "fmt"
)

/*
288. 单词的唯一缩写

一个单词的缩写需要遵循 <起始字母><中间字母数><结尾字母> 这样的格式。

以下是一些单词缩写的范例：

a) it                      --> it    (没有缩写)

     1
     ↓
b) d|o|g                   --> d1g

              1    1  1
     1---5----0----5--8
     ↓   ↓    ↓    ↓  ↓
c) i|nternationalizatio|n  --> i18n

              1
     1---5----0
     ↓   ↓    ↓
d) l|ocalizatio|n          --> l10n
假设你有一个字典和一个单词，请你判断该单词的缩写在这本字典中是否唯一。若单词的缩写在字典中没有任何 其他 单词与其缩写相同，则被称为单词的唯一缩写。

示例：

给定 dictionary = [ "deer", "door", "cake", "card" ]

isUnique("dear") -> false
isUnique("cart") -> true
isUnique("cane") -> false
isUnique("make") -> true
// 第一个字母+长度+最后一个字母 傻逼中等难度的题,还tm谷歌出的题真牛逼 服了
*/

type ValidWordAbbr struct {
    m map[string]string
}

func getAbbr(s string) string {
    if len(s) <= 2 {
        return s
    }
    return fmt.Sprintf("%c%d%c", s[0], len(s)-2, s[len(s)-1])
}

func Constructor(dictionary []string) ValidWordAbbr {
    m := map[string]string{}
    for _, d := range dictionary {
        k := getAbbr(d)
        if _, ok := m[k]; ok && m[k] != d {
            m[k] = ""
        } else {
            m[k] = d
        }
    }
    return ValidWordAbbr{m}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
    k := getAbbr(word)
    v, ok := this.m[k]
    return !ok || v == word
}

