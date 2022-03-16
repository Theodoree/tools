package _100_2199

/*
2103. 环和杆
总计有 n 个环，环的颜色可以是红、绿、蓝中的一种。这些环分布穿在 10 根编号为 0 到 9 的杆上。
给你一个长度为 2n 的字符串 rings ，表示这 n 个环在杆上的分布。rings 中每两个字符形成一个 颜色位置对 ，用于描述每个环：
第 i 对中的 第一个 字符表示第 i 个环的 颜色（'R'、'G'、'B'）。
第 i 对中的 第二个 字符表示第 i 个环的 位置，也就是位于哪根杆上（'0' 到 '9'）。
例如，"R3G2B1" 表示：共有 n == 3 个环，红色的环在编号为 3 的杆上，绿色的环在编号为 2 的杆上，蓝色的环在编号为 1 的杆上。
找出所有集齐 全部三种颜色 环的杆，并返回这种杆的数量。
*/
func countPoints(rings string) int {
	var target [10][3]bool

	for i := 0; i < len(rings); i += 2 {
		color := rings[i]
		ringNum := rings[i+1] - '0'

		switch color {
		case 'R':
			target[ringNum][0] = true
		case 'G':
			target[ringNum][1] = true
		case 'B':
			target[ringNum][2] = true

		}

	}

	var count int
	for _, v := range target {
		if v[0] && v[1] && v[2] {
			count++
		}
	}

	return count

}
