package _300_1399


/*
1376. 通知所有员工所需的时间
公司里有 n 名员工，每个员工的 ID 都是独一无二的，编号从 0 到 n - 1。公司的总负责人通过 headID 进行标识。
在 manager 数组中，每个员工都有一个直属负责人，其中 manager[i] 是第 i 名员工的直属负责人。对于总负责人，manager[headID] = -1。题目保证从属关系可以用树结构显示。
公司总负责人想要向公司所有员工通告一条紧急消息。他将会首先通知他的直属下属们，然后由这些下属通知他们的下属，直到所有的员工都得知这条紧急消息。
第 i 名员工需要 informTime[i] 分钟来通知它的所有直属下属（也就是说在 informTime[i] 分钟后，他的所有直属下属都可以开始传播这一消息）。
返回通知所有员工这一紧急消息所需要的 分钟数 。
*/
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	var buf  = make([][]int,n)
	for i := 0; i < n; i++ {
		if i == headID {
			continue
		}
		buf[manager[i]] = append(buf[manager[i]], i)
	}

	var max int
	dfs(buf, informTime, headID, 0, &max)
	return max
}
func dfs(buf [][]int, informTime []int, headId int, cur int, max *int) {
	if cur > *max {
		*max = cur
	}
	if len(buf[headId]) == 0 {
		return
	}
	for _, v := range buf[headId] {
		dfs(buf, informTime, v, cur+informTime[headId], max)
	}
}

