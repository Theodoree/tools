package _200_1299

/*
1291. 顺次数

我们定义「顺次数」为：每一位上的数字都比前一位上的数字大 1 的整数。

请你返回由 [low, high] 范围内所有顺次数组成的 有序 列表（从小到大排序）。



示例 1：

输出：low = 100, high = 300
输出：[123,234]
示例 2：

输出：low = 1000, high = 13000
输出：[1234,2345,3456,4567,5678,6789,12345]


提示：
10 <= low <= high <= 10^9
*/

var M = map[int]int{
    1: 2,
    2: 3,
    3: 4,
    4: 5,
    5: 6,
    6: 7,
    7: 8,
    8: 9,
}

func sequentialDigits(low int, high int) []int {
    if low > high {
        return nil
    }
    
    var (
        array   []int
        curArr  []int
        cnt     int
        lowlTop int
        lowArr  []int
    )
    for low > 0 {
        lowlTop = low % 10
        lowArr = append([]int{lowlTop}, lowArr...)
        low /= 10
    }
    
    curLow := len(lowArr) - 1
    curArr = append(curArr, lowlTop)
    
    for curLow > 0 {
        curArr = append(curArr, M[curArr[len(curArr)-1]])
        curLow--
    }
    
    cnt = GetCnt(curArr)
    low = GetCnt(lowArr)
    if cnt < low {
        for i := 0; i < len(curArr); i++ {
            curArr[i] = M[curArr[i]]
        }
        cnt = GetCnt(curArr)
    }
    
    if curArr[len(curArr)-1] == 0 {
        for i := 0; i < len(curArr); i++ {
            curArr[i] = i + 1
        }
        curArr = append(curArr, len(curArr)+1)
        cnt = GetCnt(curArr)
    }
    
    
    if cnt > high || curArr[len(curArr)-1] == 0 {
        return nil
    }
    
    array = append(array, cnt)
    
    for {
        if curArr[len(curArr)-1] == 9 {
            for i := 0; i < len(curArr); i++ {
                curArr[i] = i + 1
            }
            curArr = append(curArr, len(curArr)+1)
        } else {
            for i := 0; i < len(curArr); i++ {
                curArr[i] = M[curArr[i]]
            }
        }
        cnt = GetCnt(curArr)
        if cnt > high {
            break
        }
        
        array = append(array, cnt)
    }
    
    return array
}

func GetCnt(i []int) (cnt int) {
    for _, v := range i {
        cnt *= 10
        cnt += v
    }
    
    return
}
