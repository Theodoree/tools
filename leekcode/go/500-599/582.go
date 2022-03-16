package _00_599



/*
582. 杀掉进程
系统中存在 n 个进程，形成一个有根树结构。给你两个整数数组 pid 和 ppid ，其中 pid[i] 是第 i 个进程的 ID ，ppid[i] 是第 i 个进程的父进程 ID 。
每一个进程只有 一个父进程 ，但是可能会有 一个或者多个子进程 。只有一个进程的 ppid[i] = 0 ，意味着这个进程 没有父进程 。
当一个进程 被杀掉 的时候，它所有的子进程和后代进程都要被杀掉。
给你一个整数 kill 表示要杀掉​​进程的 ID ，返回杀掉该进程后的所有进程 ID 的列表。可以按 任意顺序 返回答案。
*/
func killProcess(pid []int, ppid []int, kill int) []int {
	var m = map[int][]int{}
	for idx,v:=range pid{
		m[ppid[idx]] = append(m[ppid[idx]],v)
	}


	var arr  = []int{kill}
	var result []int
	for len(arr) >0 {
		kill = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		result = append(result,kill)
		arr = append(arr,m[kill]...)
	}
	return result

}
