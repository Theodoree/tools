package _00_299

func minMeetingRooms(intervals [][]int) int {
	var nums []int
	for _, v := range intervals {
		nums = append(nums, v[0] * 10 + 2)
		nums = append(nums, v[1] * 10 + 1)
	}
	sort.Ints(nums)
	maxRoom := 0
	curNeedRoom := 0
	for _, v := range nums {
		if v % 10 == 1 {
			curNeedRoom--
		} else {
			curNeedRoom++
		}
		if curNeedRoom > maxRoom {
			maxRoom = curNeedRoom
		}
	}
	return maxRoom
}
