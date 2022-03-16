package _100_2199


/*
2121. 相同元素的间隔之和
给你一个下标从 0 开始、由 n 个整数组成的数组 arr 。
arr 中两个元素的 间隔 定义为它们下标之间的 绝对差 。更正式地，arr[i] 和 arr[j] 之间的间隔是 |i - j| 。
返回一个长度为 n 的数组 intervals ，其中 intervals[i] 是 arr[i] 和 arr 中每个相同元素（与 arr[i] 的值相同）的 间隔之和 。
注意：|x| 是 x 的绝对值。
*/

func getDistances(arr []int) []int64 {
	var m  = make(map[int][]int)
	for idx,v:=range arr{
		_arr:=m[v]
		if len(_arr) > 0 {
			m[v] = append(m[v],idx,idx +_arr[len(_arr)-1])
		}else{
			m[v] = append(m[v],idx,idx)
		}
	}


	result := make([]int64,len(arr))

	for key,_arr:=range m {
		for i:=0;i<len(_arr)/2;i++{
			idx:=_arr[i*2]
			val:=_arr[i*2+1]
			if i == 0 {
				result[idx]	= int64(_arr[len(_arr)-1]- idx*len(_arr)/2)
			}else if i  == (len(_arr)/2-1) {
				result[idx]	= int64(idx*len(_arr)/2 - _arr[len(_arr)-1])
			}	else {
				_ = key
				result[idx]= int64(idx*i - _arr[i*2-1])
				result[idx] += int64(_arr[len(_arr)-1]-val- idx*(len(_arr)/2-i-1))
			}
		}
	}

	return result
}

