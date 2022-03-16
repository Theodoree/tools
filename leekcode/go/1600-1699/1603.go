package _600_1699

/*
1603. 设计停车系统
请你给一个停车场设计一个停车系统。停车场总共有三种不同大小的车位：大，中和小，每种尺寸分别有固定数目的车位。
请你实现 ParkingSystem 类：
ParkingSystem(int big, int medium, int small) 初始化 ParkingSystem 类，三个参数分别对应每种停车位的数目。
bool addCar(int carType) 检查是否有 carType 对应的停车位。 carType 有三种类型：大，中，小，分别用数字 1， 2 和 3 表示。一辆车只能停在  carType 对应尺寸的停车位中。如果没有空车位，请返回 false ，否则将该车停入车位并返回 true 。
*/

type ParkingSystem struct {
	park [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	var p ParkingSystem
	p.park[0] = big
	p.park[1] = medium
	p.park[2] = small
	return p
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.park[carType-1] == 0 {
		return false
	}
	this.park[carType-1]--
	return true
}
