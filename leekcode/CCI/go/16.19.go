package _go

import "sort"

/*
面试题 16.19. 水域大小
你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
*/
func pondSizes(land [][]int) []int {
	var result []int
	for i:=0;i<len(land);i++{

		for j:=0;j<len(land[i]);j++{
			if land[i][j] == 0 {
				cnt:= bfs(land,i,j)
				result  = append(result,cnt)
			}
		}

	}
	sort.Ints(result)
	return result
}


func bfs(land [][]int,i,j int) int {
	var cnt int

	idxs:=[]int{i,j}
	land[i][j] = -1
	for len(idxs) > 0 {
		i = idxs[len(idxs)-2]
		j = idxs[len(idxs)-1]
		idxs =idxs [:len(idxs)-2]
		cnt++


		// 上
		if i-1 >=0 && land[i-1][j] == 0 {
			idxs = append(idxs,i-1,j)
			land[i-1][j] = -1
		}

		// 下
		if i+1 < len(land) && land[i+1][j] == 0 {
			idxs = append(idxs,i+1,j)
			land[i+1][j] = -1
		}

		// 左
		if j-1 >=0 && land[i][j-1] == 0 {
			idxs = append(idxs,i,j-1)
			land[i][j-1] = -1
		}

		// 右
		if j+1 < len(land[0]) && land[i][j+1] == 0 {
			idxs = append(idxs,i,j+1)
			land[i][j+1] = -1
		}


		// 左上
		if i-1 >=0 && j-1 >=0 && land[i-1][j-1] == 0 {
			idxs = append(idxs,i-1,j-1)
			land[i-1][j-1] = -1
		}

		// 右上
		if i-1 >=0 && j+1 < len(land[0]) && land[i-1][j+1] == 0 {
			idxs = append(idxs,i-1,j+1)
			land[i-1][j+1] = -1
		}


		// 左下
		if i+1 < len(land) && j-1 >=0 && land[i+1][j-1] == 0 {
			idxs = append(idxs,i+1,j-1)
			land[i+1][j-1] = -1
		}


		// 右下
		if i+1 < len(land) && j+1 < len(land[0]) && land[i+1][j+1] == 0 {
			idxs = append(idxs,i+1,j+1)
			land[i+1][j+1] = -1
		}
	}
	return cnt
}
