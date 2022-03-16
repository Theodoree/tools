package _500_1599


/*
1560. 圆形赛道上经过次数最多的扇区
给你一个整数 n 和一个整数数组 rounds 。有一条圆形赛道由 n 个扇区组成，扇区编号从 1 到 n 。现将在这条赛道上举办一场马拉松比赛，该马拉松全程由 m 个阶段组成。其中，第 i 个阶段将会从扇区 rounds[i - 1] 开始，到扇区 rounds[i] 结束。举例来说，第 1 阶段从 rounds[0] 开始，到 rounds[1] 结束。
请你以数组形式返回经过次数最多的那几个扇区，按扇区编号 升序 排列。
注意，赛道按扇区编号升序逆时针形成一个圆（请参见第一个示例）。
*/

func mostVisited(n int, rounds []int) []int {
	var buf = make([]int,n)
	buf[rounds[0]-1]++
	for i:=1;i<len(rounds);i++{
		cur:=rounds[i]
		if cur <= rounds[i-1]{
			cur+=n
		}
		for j:=rounds[i-1]+1;j<=cur;j++{
			buf[(j-1)%len(buf)]++
		}
	}

	var max int
	for _,v:=range buf{
		if v > max {
			max = v
		}
	}

	var result []int
	for idx,v:=range buf{
		if v == max {
			result = append(result,idx+1)
		}
	}

	return result

}

