package _00_799



/*
787. K 站中转内最便宜的航班
有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，表示该航班都从城市 fromi 开始，以价格 pricei 抵达 toi。
现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的 价格最便宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
*/

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	maps := make([][]int,n)

	for i:=0;i<len(maps);i++{
		maps[i] = make([]int,n)
	}

	for _,v:=range flights{
		maps[v[1]][v[0]] = v[2]
		// to -> from[0] from[1] from[2] ...
	}


	min:=math.MaxInt64
	dfs(dst,maps,src,k,0,&min)
	if min ==math.MaxInt64{
		return -1
	}
	return min
}


func dfs(city int,flights [][]int,src int,k int,price int,mostCheap *int){
	if city == src &&  *mostCheap > price{
		*mostCheap = price
		return
	}
	if k < 0 {
		return
	}
	for _city,p:=range flights[city]{
		if (price == 0 && _city != src) || p+price >= *mostCheap {
			continue
		}
		dfs(_city,flights,src,k-1,p+price,mostCheap)
	}
}

func min(a int,b int) int {
	if a < b {
		return a
	}
	return b
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	min_val := 10000*101+1
	m1 := make([][]int,k+2)
	for i:= range m1 {
		m1[i] = make([]int,n)
		for j:= range m1[i]{
			m1[i][j] = min_val
		}
	}
	m1[0][src] = 0
	for s:=1;s<k+2;s++ {
		for _,ran := range flights{
			from,to,cost := ran[0],ran[1],ran[2]
			m1[s][to] = min(m1[s][to],m1[s-1][from] + cost)
		}
	}
	var get_min int = min_val
	for s:=1;s<k+2;s++ {
		get_min = min(get_min,m1[s][dst])
	}

	if (get_min == min_val) {
		return -1
	}

	return get_min
}


