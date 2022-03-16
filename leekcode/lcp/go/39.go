package _go

/*
LCP 39. 无人机方阵
在 「力扣挑战赛」 开幕式的压轴节目 「无人机方阵」中，每一架无人机展示一种灯光颜色。 无人机方阵通过两种操作进行颜色图案变换：
调整无人机的位置布局
切换无人机展示的灯光颜色
给定两个大小均为 N*M 的二维数组 source 和 target 表示无人机方阵表演的两种颜色图案，由于无人机切换灯光颜色的耗能很大，请返回从 source 到 target 最少需要多少架无人机切换灯光颜色。
注意： 调整无人机的位置布局时无人机的位置可以随意变动。
示例 1：
输入：source = [[1,3],[5,4]], target = [[3,1],[6,5]]
*/
func minimumSwitchingTimes(source [][]int, target [][]int) int {
	m, n := len(source), len(source[0])
	var cnt [10001]int16

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt[source[i][j]]++
			cnt[target[i][j]]--
		}
	}

	differentCount := 0

	abs := func(i int16) int {
		if i < 0 {
			i *= -1
		}
		return int(i)
	}

	for _, c := range cnt {
		if c == 0 {
			continue
		}
		differentCount += abs(c)
	}

	return differentCount / 2
}
