package __99

import (
    "container/list"
    "math"
)

/*
42. 接雨水

提交记录
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/
// 动态划分
func trap(height []int) int {
    
    // 动态划分
    left := make([]int, len(height))
    right := make([]int, len(height))
    
    if len(height) > 0 {
        left[0] = height[0]
    }
    for i := 1; i < len(height); i++ {
        left[i] = max(left[i-1], height[i])
    }
    
    for i := len(height) - 2; i >= 0; i-- {
        right[i] = max(right[i+1], height[i+1])
    }
    
    var water int
    for i := 0; i < len(height); i++ {
        water += max(0, min(left[i], right[i])-height[i])
    }
    
    return water
}

// 链表
func trap(height []int) int {
    var sum, shortHigh, distance int // sum为装水总量
    stack := list.New()              // 注意stack用作栈
    var top *list.Element            // 栈顶元素节点， 其内存储数组下标
    current := 0                     // 指针
    for current < len(height) {
        
        // 栈非空，且栈顶元素所指向的高度 < 当前高度
        for stack.Len() != 0 && height[current] > height[stack.Back().Value.(int)] {
            top = stack.Back() // 获取栈顶元素的指针
            stack.Remove(top)  // 移除栈顶元素
            
            if stack.Len() == 0 {
                break
            } // 栈空则退出内层循环
            distance = current - stack.Back().Value.(int) - 1 // 两堵墙之间的距离
            shortHigh = int(math.Min(float64(height[stack.Back().Value.(int)]), float64(height[current])))
            sum = sum + distance*(shortHigh-height[top.Value.(int)])
        }
        // 将当前指向的墙的下标压入栈，并将current后移
        stack.PushBack(current)
        current++
    }
    
    return sum
}
func min(a, b int) int {
    if a > b {
        return b
    }
    return a
    
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
