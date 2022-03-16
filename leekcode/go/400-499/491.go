package _00_499



func findSubsequences(nums []int) [][]int {
	var m = make(map[string]struct{})
	var result [][]int
	var str []string
	for _, v := range nums {
		str = append(str, strconv.Itoa(v))
	}
	backtracking(nums, str, 0, nil, "", m, &result)
	return result
}

func deepcopy(src []int) []int {
	dist := make([]int, len(src))
	copy(dist, src)
	return dist
}

func backtracking(nums []int, str []string, idx int, arr []int, cur string, m map[string]struct{}, result *[][]int) {
	if len(arr) >= 2 {
		if _, ok := m[cur]; !ok {
			*result = append(*result, arr)
			m[cur] = struct{}{}
		}
	}
	if idx >= len(nums) {
		return
	}

	if len(arr) == 0 || arr[len(arr)-1] <= nums[idx] {
		buf := deepcopy(arr)
		buf = append(buf, nums[idx])
		backtracking(nums, str, idx+1, buf, cur+" "+str[idx], m, result)
	}
	buf1 := deepcopy(arr)
	backtracking(nums, str, idx+1, buf1, cur, m, result)

}

func findSubsequences(nums []int) [][]int {
	var subRes []int
	var res [][]int
	backTring(0,nums,subRes,&res)
	return res
}
func backTring(startIndex int,nums,subRes []int,res *[][]int){
	if len(subRes)>1{
		tmp:=make([]int,len(subRes))
		copy(tmp,subRes)
		*res=append(*res,tmp)
	}
	history:=[201]int{}//记录本层元素使用记录
	for i:=startIndex;i<len(nums);i++{
		//分两种情况判断：一，当前取的元素小于子集的最后一个元素，则继续寻找下一个适合的元素
		//                或者二，当前取的元素在本层已经出现过了，所以跳过该元素，继续寻找
		if len(subRes)>0&&nums[i]<subRes[len(subRes)-1]||history[nums[i] + 100]==1{
			continue
		}
		history[nums[i] + 100]=1//表示本层该元素使用过了
		subRes=append(subRes,nums[i])
		backTring(i+1,nums,subRes,res)
		subRes=subRes[:len(subRes)-1]
	}
}

