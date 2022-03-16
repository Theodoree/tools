package _00_299

import (
	"encoding/binary"
	"strings"
)

/*
271. 字符串的编码与解码
请你设计一个算法，可以将一个 字符串列表 编码成为一个 字符串。这个编码后的字符串是可以通过网络进行高效传送的，并且可以在接收端被解码回原来的字符串列表。
1 号机（发送方）有如下函数：
*/
type Codec struct {

}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var s strings.Builder
	for i:=0;i<len(strs);i++{
		var b [4]byte
		binary.BigEndian.PutUint32(b[:],uint32(len(strs[i])))
		s.Write(b[:])
		s.WriteString(strs[i])
	}
	return s.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var arr []string
	var idx int
	for idx < len(strs) {
		l := binary.BigEndian.Uint32([]byte(strs[idx : idx+4]))
		arr = append(arr, strs[idx+4:idx+4+int(l)])
		idx += 4 + int(l)
	}
	return arr
}


