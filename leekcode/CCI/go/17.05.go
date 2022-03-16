package _go



/*
面试题 17.05.  字母与数字
给定一个放有字母和数字的数组，找到最长的子数组，且包含的字母和数字的个数相同。
返回该子数组，若存在多个最长子数组，返回左端点下标值最小的子数组。若不存在这样的数组，返回一个空数组。
*/
func findLongestSubarray(array []string) []string {
	maxLen := 0
	maxIndexChar := 0
	valueArr := make([]int, len(array)+1)
	firstIndex := make(map[int]int, 0)
	for index, value := range array {
		if []byte(value)[0] < 58 {
			valueArr[index+1] = -1
		} else {
			valueArr[index+1] = 1
		}
	}
	firstIndex[0] = 0
	for i := 1; i < len(valueArr); i++ {
		valueArr[i] = valueArr[i] + valueArr[i-1]
		start, ok := firstIndex[valueArr[i]]
		if ok == false {
			firstIndex[valueArr[i]] = i
		} else {
			if i - start > maxLen {
				maxLen = i - start
				maxIndexChar = valueArr[i]
			}
		}
	}
	if maxLen == 0 {
		return array[0:0]
	}
	return array[firstIndex[maxIndexChar]:firstIndex[maxIndexChar]+maxLen]
}
