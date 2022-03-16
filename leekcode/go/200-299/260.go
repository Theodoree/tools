package _00_299



/*
260. 只出现一次的数字 III

给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

示例 :

输入: [1,2,1,3,2,5]
输出: [3,5]
注意：

结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现
*/
func singleNumber(nums []int) []int {

    recordMap := make(map[int]string, 0)
    for _, num := range nums {
        value := recordMap[num]
        if value != "" {
            delete(recordMap, num)
        } else {
            recordMap[num] = "o"
        }
    }
    results := make([]int, 0, 0)
    for k, _ := range recordMap {
        results = append(results, k)
    }
    return results
}

func singleNumber(nums []int) []int {

    var m = make(map[int]int)
    for _, v := range nums {
        m[v]++
    }
    var cnt []int
    for key,v:=range m{
        if v == 1 {
            cnt = append(cnt,key)
        }
    }
    return  cnt

}
