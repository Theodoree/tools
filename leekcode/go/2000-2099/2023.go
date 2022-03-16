package _000_2099


/*
2023. 连接后等于目标字符串的字符串对
给你一个 数字 字符串数组 nums 和一个 数字 字符串 target ，请你返回 nums[i] + nums[j] （两个字符串连接）结果等于 target 的下标 (i, j) （需满足 i != j）的数目。
*/
func numOfPairs(nums []string, target string) int {
	var cnt int
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums);j++{
			if i == j {
				continue
			}
			if len(nums[i]) + len(nums[j]) == len(target) &&
				nums[i] == target[:len(nums[i])] &&
				nums[j] == target[len(nums[i]):] {
				cnt++
			}

		}
	}
	return cnt
}
