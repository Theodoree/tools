package _00_299



//242.有效的字母异位词

func isAnagram(s string, t string) bool {
    countsS := make([]int, 26)
    countsT := make([]int, 26)

    if s == t {
        return  true
    }
    if s == "" || t == ""{
        return  false
    }

    if len(s) != len(t) {
        return false
    }
    index := 0

    for {

        countsS[s[index]-'a']++
        countsT[t[index]-'a']++
        index++
        if index >= len(s) || index > len(t) {
            break
        }
    }

    for i:=0;i<len(countsS);i++{
        if countsS[i] != countsT[i]{
            return false
        }
    }

    return true
}
