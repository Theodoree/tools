package _700_1799

/*
1736. 替换隐藏数字得到的最晚时间
给你一个字符串 time ，格式为 hh:mm（小时：分钟），其中某几位数字被隐藏（用 ? 表示）。
有效的时间为 00:00 到 23:59 之间的所有时间，包括 00:00 和 23:59 。
替换 time 中隐藏的数字，返回你可以得到的最晚有效时间。
*/
func maximumTime(time string) string {
	buf := []byte(time)
	var flag bool
	for i := 0; i < len(time); i++ {
		switch time[i] {
		case ':':
			flag = true
		case '?':
			if flag {
				switch i {
				case 3:
					buf[i] = '5'
				case 4:
					buf[i] = '9'
				}
				continue
			}

			switch i {
			case 0: //
				if buf[i+1] > '3' && buf[i+1] != '?' {
					buf[i] = '1'
				} else {
					buf[i] = '2'
				}
			case 1:
				if buf[i-1] == '2' {
					buf[i] = '3'
				} else {
					buf[i] = '9'
				}
			}
		}

	}
	return string(buf)
}
