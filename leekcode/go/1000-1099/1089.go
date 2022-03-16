package _000_1099

import "fmt"

/*
1089. 复写零

给你一个长度固定的整数数组 arr，请你将该数组中出现的每个零都复写一遍，并将其余的元素向右平移。

注意：请不要在超过该数组长度的位置写入元素。

要求：请对输入的数组 就地 进行上述修改，不要从函数返回任何东西。

*/


func duplicateZeros(arr []int)  {

    for i:=0;i<len(arr);i++{
        if arr[i] == 0 {
            moveslice(arr,i)
            i++
        }

    }

    fmt.Println(arr)
}

func duplicateZeros1(arr []int) {
    n := len(arr)
    i, j := 0, 0
    for i = 0; i < n; i, j = i+1, j+1 {
        if arr[i] == 0 {
            j++
        }
    }
    for i = n - 1; i >= 0; i-- {
        j--
        if j < n {
            arr[j] = arr[i]
        }
        if arr[i] == 0 {
            j--
            if j < n {
                arr[j] = 0
            }
        }
    }
}

func moveslice(arr []int,k int){
    for i:=len(arr)-1;i>k;i--{
        arr[i] = arr[i-1]
    }
}