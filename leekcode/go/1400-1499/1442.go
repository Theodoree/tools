package _400_1499



/*
1442. 形成两个异或相等数组的三元组数目
给你一个整数数组 arr 。
现需要从数组中取三个下标 i、j 和 k ，其中 (0 <= i < j <= k < arr.length) 。
a 和 b 定义如下：
a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
注意：^ 表示 按位异或 操作。
请返回能够令 a == b 成立的三元组 (i, j , k) 的数目。
*/
func countTriplets(arr []int) int {

	var count int

	for i:=0;i<len(arr)-1;i++{
		sum:=0
		for k:=i;k<len(arr);k++{
			sum^=arr[k]
			if sum == 0 && k > i {
				count +=(k-i)
			}
		}
	}

	return count
}

