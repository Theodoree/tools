package _00_299



/*
210. 课程表 II
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。


*/
func findOrder(numCourses int, prerequisites [][]int) []int {

	var m = make(map[int][]int)
	for _, v := range prerequisites {
		m[v[0]] = append(m[v[0]], v[1])
	}

	var hashTable = make(map[int]int, numCourses)
	for i := 0; i < numCourses; i++ {
		exits := make(map[int]bool)
		if !dfs(m, i, i, exits, hashTable) {
			return nil
		}
		if _, ok := hashTable[i]; !ok {
			hashTable[i] = len(hashTable)
		}
	}
	var arr = make([]int, numCourses)
	for key, idx := range hashTable {
		arr[idx] = key
	}

	return arr
}

func dfs(m map[int][]int, start int, cur int, exits map[int]bool, hashTable map[int]int) bool {
	for _, v := range m[cur] {
		if v < start {
			continue
		}
		if exits[v] {
			return false
		}
		exits[v] = true
		flag := dfs(m, start, v, exits, hashTable)
		if !flag {
			return false
		}
		if _, ok := hashTable[v]; !ok {
			hashTable[v] = len(hashTable)
		}
		exits[v] = false
	}
	return true

}

