package _00_199

import "fmt"

// 杨辉三角

func generate(numRows int) [][]int {
    
    var intArray [][]int
    
    for i := 0; i < numRows; i++ {
        array := make([]int, i+1)
        array[0] = 1
        array[len(array)-1] = 1
        intArray = append(intArray, array)
    }
    
    for k, array := range intArray {
        for k1, v := range array {
            if v == 0 {
                array[k1] = intArray[k-1][k1-1] + intArray[k-1][k1]
            }
        }
    }
    
    return intArray
    
}

func main() {
    
    fmt.Println(generate(10))
    
}
