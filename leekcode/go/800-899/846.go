package _00_899



// 846
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize > 0 {
		return false
	}
	sort.Ints(hand)
	cnt := map[int]int{}
	for _, num := range hand {
		cnt[num]++
	}
	for _, x := range hand {
		if cnt[x] == 0 {
			continue
		}
		min := cnt[x]
		cnt[x] = 0

		for num := x + 1; num < x+groupSize; num++ {
			if cnt[num]-min < 0 {
				return false
			}
			cnt[num] -= min
		}
	}
	return true
}