package _00_399

/*
370. 区间加法
假设你有一个长度为 n 的数组，初始情况下所有的数字均为 0，你将会被给出 k​​​​​​​ 个更新的操作。

其中，每个操作会被表示为一个三元组：[startIndex, endIndex, inc]，你需要将子数组 A[startIndex ... endIndex]（包括 startIndex 和 endIndex）增加 inc。

请你返回 k 次操作后的数组。
*/
func getModifiedArray(length int, updates [][]int) []int {
	ans := make([]int, length+1)

	for i := range updates {
		start := updates[i][0]
		end := updates[i][1]
		inc := updates[i][2]
		ans[start] += inc
		ans[end+1] -= inc
	}
	fmt.Println(ans)
	for i := 0; i < length-1; i++ {
		ans[i+1] = ans[i] + ans[i+1]
	}

	return ans[:len(ans)-1]
}
