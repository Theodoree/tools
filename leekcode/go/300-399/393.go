package _00_399


/*
393. UTF-8 编码验证
UTF-8 中的一个字符可能的长度为 1 到 4 字节，遵循以下的规则：

对于 1 字节的字符，字节的第一位设为0，后面7位为这个符号的unicode码。
对于 n 字节的字符 (n > 1)，第一个字节的前 n 位都设为1，第 n+1 位设为0，后面字节的前两位一律设为10。剩下的没有提及的二进制位，全部为这个符号的unicode码。
这是 UTF-8 编码的工作方式：

   Char. number range  |        UTF-8 octet sequence
      (hexadecimal)    |              (binary)
   --------------------+---------------------------------------------
   0000 0000-0000 007F | 0xxxxxxx
   0000 0080-0000 07FF | 110xxxxx 10xxxxxx
   0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
   0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
给定一个表示数据的整数数组，返回它是否为有效的 utf-8 编码。

前 3 位都是 1 ，第 4 位为 0 表示它是一个3字节字符。
下一个字节是开头为 10 的延续字节，这是正确的。
但第二个延续字节不以 10 开头，所以是不符合规则的。
*/

func validUtf8(data []int) bool {
    if len(data) == 0 {
        return false
    }

    for len(data) > 0 {
        // 可能全是一字节的utf-8
        for checkOneByte(data[0]) {
            data = data[1:]
            if len(data) == 0 {
                return true
            }
        }

        l := check(data[0])
        if l == -1 || len(data) < l {
            return false
        }
        data = data[1:]
        l--

        for l > 0 {
            val := data[0]

            // 高位必须为 1 0
            if val&(1<<7) != 1<<7 || val&(1<<6) == 1<<6 {
                return false
            }

            data = data[1:]
            l--
        }

    }

    return true

}

func check(num int) int {
    switch {
    case num&0b11110000 == 0b11110000:
        if num&(1<<3) == 0 {
            return 4
        }
    case num&0b11100000 == 0b11100000:
        if num&(1<<4) == 0 {
            return 3
        }
    case num&0b11000000 == 0b11000000:
        if num&(1<<5) == 0 {
            return 2
        }
    }

    return -1
}

func checkOneByte(num int) bool {
    return num&(1<<7) == 0
}

