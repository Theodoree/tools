package _00_899

//832. 翻转图像

func flipAndInvertImage(A [][]int) [][]int {

	for i := 0; i < len(A); i++ {
		var tmp []int
		for k := len(A[i])-1; k >= 0; k-- {

			switch A[i][k] {
			case 0:
				A[i][k] = 1
			case 1:
				A[i][k] = 0
			}
			tmp = append(tmp,A[i][k])
		}
		A[i] = tmp
	}

	return A
}

func main() {

}
