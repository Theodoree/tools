package sort


func ShellSort(slice []int) []int {
	h := 1
	for h <= len(slice)/3 {
		h *= 3 + 1
	}
	for h > 0 {
		for i := h; i < len(slice); i++ {
			for j := i; j >= h && slice[j] < slice[j-h]; j -= h {
				slice[j], slice[j-h] = slice[j-h], slice[j]
			}
		}
		h = (h-1)/ 3
	}

	return slice
}
