package _200_1299



/*
1282. 用户分组
有 n 位用户参加活动，他们的 ID 从 0 到 n - 1，每位用户都 恰好 属于某一用户组。给你一个长度为 n 的数组 groupSizes，其中包含每位用户所处的用户组的大小，请你返回用户分组情况（存在的用户组以及每个组中用户的 ID）。
你可以任何顺序返回解决方案，ID 的顺序也不受限制。此外，题目给出的数据保证至少存在一种解决方案。
*/
func groupThePeople(groupSizes []int) [][]int {


	var arr = make([] [][]int,len(groupSizes)+1)


	for i:=0;i<len(groupSizes);i++{
		_cap:=groupSizes[i]
		l:=len(arr[_cap])
		if l == 0 || len(arr[_cap][l-1]) == _cap{
			subArr:=make([]int,0,_cap)
			subArr = append(subArr,i)
			arr[_cap] = append(arr[_cap],subArr)
		}else{
			arr[_cap][l-1] = append(arr[_cap][l-1],i)
		}
	}


	var result [][]int


	for i:=1;i<len(arr);i++{
		if len(arr[i]) != 0 {
			result = append(result,arr[i]...)
		}
	}

	return result
}

func groupThePeople(groupSizes []int) [][]int {
	if len(groupSizes) == 0 {
		return nil
	}

	for idx := range groupSizes {
		groupSizes[idx] <<= 16
		groupSizes[idx] |= idx
	}
	sort.Ints(groupSizes)

	var res [][]int
	curRes := []int{groupSizes[0]&0xFFFF}
	for idx := 1; idx < len(groupSizes); idx++ {
		if len(curRes) == (groupSizes[idx-1]>>16) || (groupSizes[idx]>>16) != (groupSizes[idx-1]>>16) {
			res = append(res, curRes)
			curRes = nil
		}

		curRes = append(curRes, groupSizes[idx]&0xFFFF)
	}

	if len(curRes) > 0 {
		res = append(res, curRes)
	}

	return res
}

