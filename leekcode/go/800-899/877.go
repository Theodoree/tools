package _00_899


/*
877:石子游戏
亚历克斯和李用几堆石子在做游戏。偶数堆石子排成一行，每堆都有正整数颗石子 piles[i] 。

游戏以谁手中的石子最多来决出胜负。石子的总数是奇数，所以没有平局。

亚历克斯和李轮流进行，亚历克斯先开始。 每回合，玩家从行的开始或结束处取走整堆石头。 这种情况一直持续到没有更多的石子堆为止，此时手中石子最多的玩家获胜。

假设亚历克斯和李都发挥出最佳水平，当亚历克斯赢得比赛时返回 true ，当李赢得比赛时返回 false 。

*/

func stoneGame(piles []int) bool {
    /*
       先手者必胜
       return true
    */

    /*数组方案
      if len(piles) <= 2 {
          return true
      }

      alex := 0
      lee := 0

      if piles[0] > piles[len(piles)-1] {
          alex = piles[0]
          lee = piles[len(piles)-1]
      } else {
          alex = piles[len(piles)-1]
          lee = piles[0]
      }

      left := 1
      right := len(piles) - 2
      for left < right {

          if piles[left] > piles[right] {
              alex += piles[left]
              lee += piles[right]
          } else {
              alex += piles[right]
              lee += piles[left]
          }
          left++
          right--
      }

      fmt.Println(alex,lee)
      return alex > lee
    */
    /*动态划分方案*/

    // 动态规划，二维数组
    /*
       // dp[i][j] 表示下标 i 到 j 的数堆石子中，当前玩家相对对手能获得多出的石子数，大于 0 表示胜利
       length := len(piles)
       dp := make([][]int, length)
       for i := 0; i < length; i++ {
       	dp[i] = make([]int, length)
       // 	// 初始状态，只有一堆时，多出石子数就是 piles[i]
       	dp[i][i] = piles[i]
       }

       // // j 比 i 大，最大能取最后一堆石子 length-1
       // // 因此 i 计算从 0 到 length-2 即可
       for i := 0; i < length-1; i++ {
       	for j := i + 1; j < length; j++ {
       // 		// DP 状态转移方程，
       // 		// 如果选择 i 堆石头，则 dp[i+1][j] 是对手的多出石头数，当前选手多出的石子数为 piles[i] - dp[i+1][j]
       // 		// 同理选 j 堆则是 piles[j] - dp[i][j-1]
       		dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
       	}
       }
       return dp[0][length-1] > 0
    */
    // 动态规划，优化，一维数组
    // 每个 j 值对应的一维数组可以覆盖 j-1 对应的一维数组，可以简化用一维数据记录结果
    // 该一维数组 dp[i] 表示 d[i][length-1] 的结果，即 i 开始到 length-1 的优胜石子数

    length := len(piles)
    dp := piles[:]
    for len := 2; len <= length; len++ {
        for i := 0; i < length-len+1; i++ {
            j := i + len - 1
            dp[i] = max(piles[i]-dp[i+1], piles[j]-dp[i])
        }
    }
    return dp[0] > 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
