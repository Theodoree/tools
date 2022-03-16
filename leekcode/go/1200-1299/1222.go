package _200_1299



func queensAttacktheKing(queens [][]int, king []int) [][]int {
	/*
	   1.用一个二维数组，模拟queen的分布 queen的位置置1，其余位置置0；
	   2.8个方向；
	   3.从King的位置开始，遍历这8个方向；
	   4.每选定一个方向时，一直搜
	       （1）遇到一个queen，记录下来，break 不用再沿着这个方向搜了，因为这个方向的后面即使还有，也被当前的挡住了；
	       （2）没遇到queen，沿着这个方向进行下，一直超越边界
	   5.return res
	*/
	var graphForQueens =make([][]int,8)
	for i:=range graphForQueens{
		graphForQueens[i]=make([]int,8)
	}
	//用一个二维数组，模拟queen的分布 queen的位置置1，其余位置置0
	for i:=range queens{
		oneQueen:=queens[i]
		graphForQueens[oneQueen[0]][oneQueen[1]]=1
	}
	//遍历的8个 方向，上下左右 西北 东北 西南 东南
	var dirs=[][]int{{0,-1},{0,1},{-1,0},{1,0},{-1,-1},{1,-1},{-1,1},{1,1}}
	//从King开始遍历
	var res[][]int
	for i:=range dirs {
		oneDir:=dirs[i]
		r:=king[0]
		c:=king[1]
		curR:=r+oneDir[0]
		curC:=c+oneDir[1]
		//每选定一个方向时，一直搜
		for 0<=curR&&curR<8&&0<=curC&&curC<8{
			if graphForQueens[curR][curC]==1{
				res=append(res,[]int{curR,curC})
				break
				//遇到一个queen，记录下来，break 不用再沿着这个方向搜了，因为这个方向的后面即使还有，也被当前的挡住了
			}
			//没找到 沿着8个方向继续搜
			curR=curR+oneDir[0]
			curC=curC+oneDir[1]
		}
	}
	return res
}

