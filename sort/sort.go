package sort

import "fmt"

var g_number int

//回溯
func EightQueen(num int) {
	queens := num

	ColumnIndex := []int{}

	for i := 0; i < queens; i++ {
		ColumnIndex = append(ColumnIndex, i)
	}

	Permutation(ColumnIndex, queens, 0)
}

func Permutation(ColumnIndex []int, length int, index int) {

	if index == length {
		if Check(ColumnIndex, length) {

			g_number++
			PrintQueen(ColumnIndex, length)
		}

	} else {

		for i := index; i < length; i++ {

			temp := ColumnIndex[i];
			ColumnIndex[i] = ColumnIndex[index];
			ColumnIndex[index] = temp;
			Permutation(ColumnIndex, length, index+1);
			temp = ColumnIndex[index];
			ColumnIndex[index] = ColumnIndex[i];
			ColumnIndex[i] = temp;
		}

	}

}

func Check(ColumnIndex []int, length int) bool {

	for i := 0; i < length; i++ {

		for j := i + 1; j < length; j++ {
			if (i-j == ColumnIndex[i]-ColumnIndex[j]) || (j-i == ColumnIndex[i]-ColumnIndex[j]) {
				return false
			}
		}

	}

	return true;
}

func PrintQueen(ColumnIndex []int, length int) {

	fmt.Printf("Solution %d\n", g_number);
	for i := 0; i < length; i++ {
		fmt.Printf("Solution %d\n", i);
	}

	fmt.Println();
}


//分治

//动态划分

//贪心

//分支界定
