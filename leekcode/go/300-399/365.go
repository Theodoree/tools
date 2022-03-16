package _00_399


/*
365. 水壶问题
有两个水壶，容量分别为 jug1Capacity 和 jug2Capacity 升。水的供应是无限的。确定是否有可能使用这两个壶准确得到 targetCapacity 升。

如果可以得到 targetCapacity 升水，最后请用以上水壶中的一或两个来盛放取得的 targetCapacity 升水。

你可以：

装满任意一个水壶
清空任意一个水壶
从一个水壶向另外一个水壶倒水，直到装满或者倒空
*/

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if targetCapacity == 0 {
		return true

	}
	if jug1Capacity+jug2Capacity < targetCapacity{
		return false
	}
	if  jug1Capacity > jug2Capacity {
		jug1Capacity,jug2Capacity = jug2Capacity,jug1Capacity
	}
	if jug1Capacity == 0 {
		return jug2Capacity == targetCapacity
	}
	for jug2Capacity%jug1Capacity != 0 {
		jug2Capacity,jug1Capacity = jug1Capacity,jug2Capacity%jug1Capacity
	}
	return targetCapacity % jug1Capacity == 0


}
