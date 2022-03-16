package _00_499



/*
454. 四数相加 II
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var m = make(map[int]int)

	for i:=0;i<len(nums1);i++{
		for j:=0;j<len(nums2);j++{
			m[nums1[i]+nums2[j]]++
		}
	}


	var count int
	for i:=0;i<len(nums3);i++{
		for j:=0;j<len(nums4);j++{
			v:=-(nums3[i]+nums4[j])
			if cnt,ok:=m[v];ok {
				count+=cnt
			}
		}
	}


	return count
}

