package _800_1899



/*
1865. 找出和为指定值的下标对
给你两个整数数组 nums1 和 nums2 ，请你实现一个支持下述两类查询的数据结构：
累加 ，将一个正整数加到 nums2 中指定下标对应元素上。
计数 ，统计满足 nums1[i] + nums2[j] 等于指定值的下标对 (i, j) 数目（0 <= i < nums1.length 且 0 <= j < nums2.length）。
实现 FindSumPairs 类：
FindSumPairs(int[] nums1, int[] nums2) 使用整数数组 nums1 和 nums2 初始化 FindSumPairs 对象。
void add(int index, int val) 将 val 加到 nums2[index] 上，即，执行 nums2[index] += val 。
int count(int tot) 返回满足 nums1[i] + nums2[j] == tot 的下标对 (i, j) 数目。
*/

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	m map[int]int

}


func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	var c FindSumPairs
	c.m = make(map[int]int)
	c.nums1 = nums1
	c.nums2 = nums2
	for _,v:=range nums2{
		c.m[v]++
	}
	return c
}


func (this *FindSumPairs) Add(index int, val int)  {
	this.m[this.nums2[index]]--
	this.nums2[index]+=val
	this.m[this.nums2[index]]++
}


func (this *FindSumPairs) Count(tot int) int {
	var count int
	for i:=0;i<len(this.nums1);i++ {
		if tot-this.nums1[i] >= 0 {
			count += this.m[tot-this.nums1[i]]
		}
	}
	return count
}

