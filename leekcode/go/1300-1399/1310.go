package _300_1399


/*
1310. 子数组异或查询
有一个正整数数组 arr，现给你一个对应的查询数组 queries，其中 queries[i] = [Li, Ri]。
对于每个查询 i，请你计算从 Li 到 Ri 的 XOR 值（即 arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。
并返回一个包含给定查询 queries 所有结果的数组。
*/
func xorQueries(arr []int, queries [][]int) []int {

	preSum:=make([]int,len(arr))

	for i:=0;i<len(preSum);i++{
		if i == 0 {
			preSum[i] = arr[i]
		}else {
			preSum[i] = preSum[i-1] ^ arr[i]
		}
	}

	var result []int
	for _,v:=range queries{
		if v[0] == v[1] {
			result = append(result,arr[v[0]])
		}else if v[0] == 0{
			result = append(result,preSum[v[1]])
		}else{
			result = append(result,preSum[v[1]]^preSum[v[0]-1])
		}
	}
	return result
}

