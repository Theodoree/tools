package _600_1699

/*
1630. 等差子数组
如果一个数列由至少两个元素组成，且每两个连续元素之间的差值都相同，那么这个序列就是 等差数列 。更正式地，数列 s 是等差数列，只需要满足：
对于每个有效的 i ， s[i+1] - s[i] == s[1] - s[0] 都成立。
例如，下面这些都是 等差数列 ：
*/

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var result []bool
	sub:=make([]int,len(nums))
	for i:=0;i<len(l);i++{
		left:=l[i]
		right:=r[i]
		if right-left <= 1 {
			result = append(result,true)
			continue
		}

		copy(sub,nums[left:right+1])
		sort.Ints(sub[:right-left+1])
		c:=sub[1]-sub[0]
		ok:=true
		for j:=2;j<=right-left;j++{
			if sub[j] - sub[j-1] != c {
				ok = false
				break
			}
		}
		result = append(result,ok)
	}

	return result
}

