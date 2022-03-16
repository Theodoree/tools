package _00_299



/*
207. 课程表
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
*/
func canFinish(numCourses int, prerequisites [][]int) bool {

	var m = make(map[int][]int)
	for _, v := range prerequisites {
		m[v[0]] = append(m[v[0]], v[1])
	}

	for i := 0; i <= numCourses; i++ {
		exits := make(map[int]bool)
		if !dfs(m, i, i, exits) {
			return false
		}
	}

	return true
}

func dfs(m map[int][]int, start int, cur int, exits map[int]bool) bool {
	if cur < start {
		return true
	}
	for _, v := range m[cur] {
		if exits[v] {
			return false
		}
		exits[v] = true
		if !dfs(m, start, v, exits) {
			return false
		}
		exits[v] = false
	}
	return true

}