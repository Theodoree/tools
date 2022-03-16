package _00_299



/*
286. 墙与门
你被给定一个 m × n 的二维网格 rooms ，网格中有以下三种可能的初始化值：
-1 表示墙或是障碍物
0 表示一扇门
INF 无限表示一个空的房间。然后，我们用 231 - 1 = 2147483647 代表 INF。你可以认为通往门的距离总是小于 2147483647 的。
你要给每个空房间位上填上该房间到 最近门的距离 ，如果无法到达门，则填 INF 即可。
*/

func wallsAndGates(rooms [][]int) {
	visit := make([][]bool, len(rooms))
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, len(rooms[0]))
	}
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			if rooms[i][j] == 0 {
				bfs(rooms, visit, []int{i, j, 0})
			}
		}
	}
}

func bfs(rooms [][]int, visit [][]bool, arr []int) {

	tmp := []int{}

	for len(arr) > 0 {
		for i := 0; i < len(arr)/3; i++ {
			r := arr[i*3]
			c := arr[i*3+1]
			cnt := arr[i*3+2]

			if visit[r][c] && rooms[r][c] <= cnt {
				continue
			}
			if rooms[r][c] != 0 {
				rooms[r][c] = cnt
			}
			visit[r][c] = true

			// 上
			if r-1 >= 0 && rooms[r-1][c] > 0 {
				tmp = append(tmp, r-1, c, cnt+1)
			}

			// 下
			if r+1 < len(rooms) && rooms[r+1][c] > 0 {
				tmp = append(tmp, r+1, c, cnt+1)
			}

			// 左
			if c-1 >= 0 && rooms[r][c-1] > 0 {
				tmp = append(tmp, r, c-1, cnt+1)
			}

			// 右
			if c+1 < len(rooms[0]) && rooms[r][c+1] > 0 {
				tmp = append(tmp, r, c+1, cnt+1)
			}

		}
		arr, tmp = tmp, arr
		tmp = tmp[:0]
	}
}

