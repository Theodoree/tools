package _00_999


/*
974. 和可被 K 整除的子数组
给定一个整数数组 nums 和一个整数 k ，返回其中元素之和可被 k 整除的（连续、非空） 子数组 的数目。
子数组 是数组的 连续 部分。
*/
func subarraysDivByK(nums []int, k int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range nums {
		sum += elem
		modulus := (sum % k + k) % k
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}
