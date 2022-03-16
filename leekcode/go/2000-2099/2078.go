package _000_2099

/*
2078. 两栋颜色不同且距离最远的房子
街上有 n 栋房子整齐地排成一列，每栋房子都粉刷上了漂亮的颜色。给你一个下标从 0 开始且长度为 n 的整数数组 colors ，其中 colors[i] 表示第  i 栋房子的颜色。
返回 两栋 颜色 不同 房子之间的 最大 距离。
第 i 栋房子和第 j 栋房子之间的距离是 abs(i - j) ，其中 abs(x) 是 x 的绝对值。
*/
func maxDistance(colors []int) int {
	var max int

	for i := 0; i < len(colors)-max; i++ {
		for j := len(colors) - 1; j > i; j-- {
			if colors[j] != colors[i] && j-i > max {
				max = j - i
			}
		}
	}
	return max
}
