package _go


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

