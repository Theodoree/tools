package _300_1399


/*
1306. 跳跃游戏 III
这里有一个非负整数数组 arr，你最开始位于该数组的起始下标 start 处。当你位于下标 i 处时，你可以跳到 i + arr[i] 或者 i - arr[i]。
请你判断自己是否能够跳到对应元素值为 0 的 任一 下标处。
注意，不管是什么情况下，你都无法跳到数组之外。
*/
func canReach(arr []int, start int) bool {
	return dfs(arr,start)
}

func dfs(arr []int, start int) bool {
	if start < 0  || start >=  len(arr) || arr[start] == -1{
		return false
	}

	step:=arr[start]
	arr[start] = -1
	return step == 0 || dfs(arr,start +step) || dfs(arr,start -step)
}

