package _00_899


/*
844. 比较含退格的字符串

给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。
*/
func backspaceCompare(S string, T string) bool {

    sByte := []byte(S)
    tByte := []byte(T)

    S = ""
    T = ""
    for i := 0; i < len(sByte); i++ {

        if sByte[i] == '#'{
            sByte[i] = ' '
            left := i -1
            for left >= 0 {
                if sByte[left] != ' '{
                    sByte[left] = ' '
                    break
                }
                left--
            }
        }
    }

    for _,v:=range sByte{
        if v != ' '{
            S+=string(v)
        }
    }

    for i := 0; i < len(tByte); i++ {

        if tByte[i] == '#'{
            tByte[i] = ' '
            left := i -1
            for left >= 0 {
                if tByte[left] != ' '{
                    tByte[left] = ' '
                    break
                }
                left--
            }
        }
    }

    for _,v:=range tByte{
        if v != ' '{
            T+=string(v)
        }
    }

    return  S == T
}
