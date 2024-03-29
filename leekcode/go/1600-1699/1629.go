package _600_1699

/*
1629. 按键持续时间最长的键
LeetCode 设计了一款新式键盘，正在测试其可用性。测试人员将会点击一系列键（总计 n 个），每次一个。
给你一个长度为 n 的字符串 keysPressed ，其中 keysPressed[i] 表示测试序列中第 i 个被按下的键。releaseTimes 是一个升序排列的列表，其中 releaseTimes[i] 表示松开第 i 个键的时间。字符串和数组的 下标都从 0 开始 。第 0 个键在时间为 0 时被按下，接下来每个键都 恰好 在前一个键松开时被按下。
测试人员想要找出按键 持续时间最长 的键。第 i 次按键的持续时间为 releaseTimes[i] - releaseTimes[i - 1] ，第 0 次按键的持续时间为 releaseTimes[0] 。
注意，测试期间，同一个键可以在不同时刻被多次按下，而每次的持续时间都可能不同。
请返回按键 持续时间最长 的键，如果有多个这样的键，则返回 按字母顺序排列最大 的那个键。
*/
func slowestKey(releaseTimes []int, keysPressed string) byte {
	var buf [26]int
	var max int
	start := 0
	for i := 0; i < len(keysPressed); i++ {
		ts := releaseTimes[i] - start
		if ts > buf[keysPressed[i]-'a'] {
			buf[keysPressed[i]-'a'] = ts
			if ts > max {
				max = ts
			}
		}
		start = releaseTimes[i]
	}

	for i := 25; i >= 0; i-- {
		if buf[i] == max {
			return byte(i + 'a')
		}
	}
	return 0

}
