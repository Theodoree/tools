package _00_199

/*
134. 加油站
在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

说明:

如果题目有解，该答案即为唯一答案。
输入数组均为非空数组，且长度相同。
输入数组中的元素均为非负数。
*/
func canCompleteCircuit(gas []int, cost []int) int {
	curIdx, curGas, count := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		val := gas[i] - cost[i]
		curGas += val
		count += val
		if curGas < 0 {
			curGas = 0
			curIdx = i + 1
		}
	}
	if count < 0 {
		return -1
	}
	return curIdx
}
