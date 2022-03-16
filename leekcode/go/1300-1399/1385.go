package _300_1399

/*
1385. 两个数组间的距离值
给你两个整数数组 arr1 ， arr2 和一个整数 d ，请你返回两个数组之间的 距离值 。
「距离值」 定义为符合此距离要求的元素数目：对于元素 arr1[i] ，不存在任何元素 arr2[j] 满足 |arr1[i]-arr2[j]| <= d 。
*/
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var count int
	for i := 0; i < len(arr1); i++ {

		flag := true
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				flag = false
				break
			}
		}
		if flag {
			count++
		}
	}
	return count

}
func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}
