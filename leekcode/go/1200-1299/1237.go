package _200_1299



/*
1237. 找出给定方程的正整数解
给你一个函数  f(x, y) 和一个目标结果 z，函数公式未知，请你计算方程 f(x,y) == z 所有可能的正整数 数对 x 和 y。满足条件的结果数对可以按任意顺序返回。

尽管函数的具体式子未知，但它是单调递增函数，也就是说：

f(x, y) < f(x + 1, y)
f(x, y) < f(x, y + 1)
*/
func findSolution(customFunction func(int, int) int, z int) [][]int {
	var result [][]int

	for i := 1; i <= z; i++ {
		for j := 1; j <= z; j++ {

			if customFunction(i, j) == z {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}
