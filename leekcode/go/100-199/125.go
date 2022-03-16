package _00_199

import "strings"

func isPalindrome(s string) bool {
    s=strings.ToLower(s)
    i,j:=0,len(s)-1
    for i<j{
        if !('a'<=s[i]&&s[i]<='z'||'0'<=s[i]&&s[i]<='9'){
            i++
            continue
        }
        if !('a'<=s[j]&&s[j]<='z'||'0'<=s[j]&&s[j]<='9'){
            j--
            continue
        }
        if s[i]!=s[j]{
            return false
        }
        i++
        j--
    }
    return true
}