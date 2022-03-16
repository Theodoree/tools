package __99

/*
93. 复原 IP 地址
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
*/
func restoreIpAddresses(s string) []string {
	var result []string
	backtracking(s, 0, nil, 0, &result)
	return result
}

func backtracking(s string, idx int, cur []byte, spilt int, result *[]string) {
	if spilt > 3 {
		return
	}
	if len(s) <= idx {
		if spilt == 3 {
			*result = append(*result, string(cur))
		}
		return
	}
	flag := s[idx] == '0'
	sum := 0
	for i := idx; i < idx+3 && i < len(s); i++ {
		sum = sum*10 + int(s[i]-'0')
		if sum > 255 {
			return
		}
		cur = append(cur, s[i])
		buf := make([]byte, len(cur))
		copy(buf[:], cur[:])
		if i == len(s)-1 {
			backtracking(s, i+1, buf, spilt, result)
		} else {
			buf = append(buf, '.')
			backtracking(s, i+1, buf, spilt+1, result)
		}
		if flag {
			return
		}

	}
}
