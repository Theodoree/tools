package _200_1299

/*
1266. 访问所有点的最小时间
平面上有 n 个点，点的位置用整数坐标表示 points[i] = [xi, yi] 。请你计算访问所有这些点需要的 最小时间（以秒为单位）。
你需要按照下面的规则在平面上移动：
每一秒内，你可以：
	沿水平方向移动一个单位长度，或者 0,+-1
	沿竖直方向移动一个单位长度，或者	+-1,0
	跨过对角线移动 sqrt(2) 个单位长度（可以看作在一秒内向水平和竖直方向各移动一个单位长度）。  +-1,+-1
必须按照数组中出现的顺序来访问这些点。
在访问某个点时，可以经过该点后面出现的点，但经过的那些点不算作有效访问。
*/


func minTimeToVisitAllPoints(points [][]int) int {
	min:= func(i,j int) int {
		if j< i {
			return j
		}
		return i
	}

	abs:= func(i int) int {
		if i < 0 {
			i*=-1
		}
		return i
	}
	max:= func(i,j int) int {
		if i> j {
			return i
		}
		return j
	}

	var sec int
	var startR,startC int
	startR = points[0][0]
	startC = points[0][1]
	for i:=1;i<len(points);i++{

		if points[i][0] == startR && points[i][1] == startC{
			continue
		}

		rMove:=abs(points[i][0]-startR)
		cMove:=abs(points[i][1]-startC)


		if cMove == 0 || rMove == 0 { // 有一个为零,横轴或者竖轴
			sec+=max(rMove,cMove)
		}else { // 那么都不为零
			sec+=min(rMove,cMove)
			sec+=abs(rMove-cMove)
		}




		startR = points[i][0]
		startC = points[i][1]
	}

	return sec
}

