package _800_1899


/*
1817. 查找用户活跃分钟数
给你用户在 LeetCode 的操作日志，和一个整数 k 。日志用一个二维整数数组 logs 表示，其中每个 logs[i] = [IDi, timei] 表示 ID 为 IDi 的用户在 timei 分钟时执行了某个操作。
多个用户 可以同时执行操作，单个用户可以在同一分钟内执行 多个操作 。
指定用户的 用户活跃分钟数（user active minutes，UAM） 定义为用户对 LeetCode 执行操作的 唯一分钟数 。 即使一分钟内执行多个操作，也只能按一分钟计数。
请你统计用户活跃分钟数的分布情况，统计结果是一个长度为 k 且 下标从 1 开始计数 的数组 answer ，对于每个 j（1 <= j <= k），answer[j] 表示 用户活跃分钟数 等于 j 的用户数。
返回上面描述的答案数组 answer 。
*/
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	result := make([]int, k)
	sort.Slice(logs, func(i, j int) bool{
		if logs[i][0] < logs[j][0] {
			return true
		} else if logs[i][0] > logs[j][0] {
			return false
		}
		if logs[i][1] < logs[j][1] {
			return true
		}
		return false
	})
	// fmt.Println(logs)
	prev := []int{-1, 0}
	stat := 0
	for _, log := range logs {
		if log[0] == prev[0] {
			if log[1] == prev[1] {
				continue
			} else {
				stat++
			}
		} else {
			if stat > 0 {
				result[stat-1]++
			}
			stat = 1
		}
		prev = log
	}
	if stat > 0 {
		result[stat-1]++
	}
	return result
}
