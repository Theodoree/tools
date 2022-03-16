package _700_1799


/*
1710. 卡车上的最大单元数
请你将一些箱子装在 一辆卡车 上。给你一个二维数组 boxTypes ，其中 boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi] ：
numberOfBoxesi 是类型 i 的箱子的数量。
numberOfUnitsPerBoxi 是类型 i 每个箱子可以装载的单元数量。
整数 truckSize 表示卡车上可以装载 箱子 的 最大数量 。只要箱子数量不超过 truckSize ，你就可以选择任意箱子装到卡车上。
返回卡车可以装载 单元 的 最大 总数。
*/
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	var sum int

	for truckSize > 0 && len(boxTypes) >0 {
		if truckSize >= boxTypes[0][0] {
			truckSize -= boxTypes[0][0]
			sum+=boxTypes[0][0]*boxTypes[0][1]
		}else{

			sum+=truckSize*boxTypes[0][1]
			break
		}
		boxTypes = boxTypes[1:]
	}
	return sum
}

