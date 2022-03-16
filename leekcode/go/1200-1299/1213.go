package _200_1299

import (
    "sort"
)

/*

1213. 三个有序数组的交集

给出三个均为 严格递增排列 的整数数组 arr1，arr2 和 arr3。

返回一个由 仅 在这三个数组中 同时出现 的整数所构成的有序数组。



示例：

输入: arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9], arr3 = [1,3,4,5,8]
输出: [1,5]
解释: 只有 1 和 5 同时在这三个数组中出现.


提示：

1 <= arr1.length, arr2.length, arr3.length <= 1000
1 <= arr1[i], arr2[i], arr3[i] <= 2000
*/

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    
    m := make(map[int]int)
    
    for i:=0;i<len(arr1);i++{
        m[arr1[i]]++
        m[arr2[i]]++
        m[arr3[i]]++
    }
    
    var record []int
    
    for k,v:=range m{
        if v >= 3 {
            record = append(record,k)
        }
    }
    sort.Ints(record)
    
    return  record
}
