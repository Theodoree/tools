package _100_2199


func minimumRefill(plants []int, capacityA int, capacityB int) int {
	ans := 0
	i := 0
	j := len(plants) - 1
	curA := capacityA
	curB := capacityB
	for i < j {
		if curA < plants[i] {
			curA = capacityA
			ans++
		}
		curA -= plants[i]
		i++
		if curB < plants[j] {
			curB = capacityB
			ans++
		}
		curB -= plants[j]
		j--
	}
	if i == j {
		if curA < plants[i] && curB < plants[i]{
			ans++
		}
	}
	return ans
}

