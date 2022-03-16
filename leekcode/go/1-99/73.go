package __99

import "fmt"

/*
1.原地算法
一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个常数空间的解决方案吗？
示例 1:

输入:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2:

输入:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
*/

//73. 矩阵置零
func setZeroes(matrix [][]int) {
    s := make([]bool, len(matrix[0]))
    var flag bool
    for _, array := range matrix {
        flag = false
        for i, v := range array {
            if v == 0 {
                flag = true
                s[i] = true
            }
            if flag {
                array[i] = 0
            }
        }

        if flag {
            for i, v := range array {
                if v == 0 {
                    break
                }
                array[i] = 0
            }
        }
    }

    for _, array := range matrix {
        for i, _ := range array {
          if s[i] {
              if array[i] == 0 {
                  break
              }
              array[i] = 0
          }
        }

    }
}

func main(){
    f:=[][]int{{1,1,1},
        {1,0,1},
        {1,1,1}}
    setZeroes(f)

    for _,v:=range f{
        fmt.Println(v)
    }


    f=[][]int{{0,1,2,0},
        {3,4,5,2},
        {1,3,1,5}}
    setZeroes(f)

    for _,v:=range f{
        fmt.Println(v)
    }


}



