package _100_2199

/*
2149. 按符号重排数组
给你一个下标从 0 开始的整数数组 nums ，数组长度为 偶数 ，由数目相等的正整数和负整数组成。
你需要 重排 nums 中的元素，使修改后的数组满足下述条件：

任意 连续 的两个整数 符号相反
对于符号相同的所有整数，保留 它们在 nums 中的 顺序 。
重排后数组以正整数开头。

重排元素满足上述条件后，返回修改后的数组。
*/
func rearrangeArray(nums []int) []int {
	var  arr = [2][]int{make([]int,0,len(nums)/2),make([]int,0,len(nums)/2)}


	for i:=0;i<len(nums);i++{
		if nums[i] >= 0 {
			arr[0] = append(arr[0],nums[i])
		}else{
			arr[1] = append(arr[1],nums[i])
		}
	}


	for i:=0;i<len(nums);i++{
		if i%2==0{
			nums[i] = arr[0][i/2]
		}else{
			nums[i] = arr[1][i/2]
		}
	}


	return nums
}

