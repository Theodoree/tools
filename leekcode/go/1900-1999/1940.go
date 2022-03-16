package _900_1999



/*
1940. 排序数组之间的最长公共子序列
给定一个由整数数组组成的数组arrays，其中arrays[i]是严格递增排序的，返回一个表示所有数组之间的最长公共子序列的整数数组。

子序列是从另一个序列派生出来的序列，删除一些元素或不删除任何元素，而不改变其余元素的顺序。



2 <= arrays.length <= 100
1 <= arrays[i].length <= 100
1 <= arrays[i][j] <= 100
*/
func longestCommonSubsequence(arrays [][]int) []int {

	var max int
	count:=make([][]int,len(arrays))
	for i:=0;i<len(count);i++{
		count[i] = make([]int,102)
	}

	for idx,v:=range arrays{
		for _,num:=range v {
			count[idx][num]++
			if num > max {
				max = num
			}
		}
	}

	var result []int
	for num:=1;num<102 && num <= max;num++{
		arr:=count[0]
		t:=arr[num]
		if t == 0 {
			continue
		}

		for j:=0;j<len(count) && t > 0;j++{
			if count[j][num] < t {
				t =count[j][num]
			}
		}

		for j:=0;j<t;j++{
			result = append(result,num)
		}

	}
	return result
}
