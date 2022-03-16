package _00_699

/*
681. 最近时刻
给定一个形如 "HH:MM" 表示的时刻 time ，利用当前出现过的数字构造下一个距离当前时间最近的时刻。每个出现数字都可以被无限次使用。
你可以认为给定的字符串一定是合法的。例如， "01:34" 和  "12:09" 是合法的，“1:34” 和 “12:9” 是不合法的。
*/
func nextClosestTime(time string) string {
	var cnt [10]bool
	cnt[time[0]-'0'] = true
	cnt[time[1]-'0'] = true
	cnt[time[3]-'0'] = true
	cnt[time[4]-'0'] = true

	n := make([]byte, len(time))
	copy(n, time)

	replace := func(start byte, end byte, idx int) bool {
		for i := start + 1; i <= end; i++ {
			if cnt[i-'0'] {
				n[idx] = i
				return true
			}
		}
		return false
	}

	if replace(time[4], '9', 4) && n[4] > time[4] {
		return string(n)
	}
	if replace(time[3], '5', 3) && n[3] > time[3] {
		replace('0'-1, '9', 4)
		return string(n)
	}

	if replace(time[1], '9', 1) && n[1] > time[1] {
		if !(n[0] >= '2' && n[1] >= '4') {
			replace('0'-1, '5', 3)
			replace('0'-1, '9', 4)
			return string(n)
		}
	}

	if replace(time[0], '2', 0) && n[0] > time[0] {
		replace('0'-1, '9', 1)
		replace('0'-1, '5', 3)
		replace('0'-1, '9', 4)
		return string(n)
	}

	replace('0'-1, '2', 0)
	replace('0', '9', 1)
	replace('0'-1, '5', 3)
	replace('0'-1, '9', 4)

	return string(n)
}

