package _00_299


/*
280. 摆动排序
给你一个的整数数组 nums, 将该数组重新排序后使 nums[0] <= nums[1] >= nums[2] <= nums[3]...
输入数组总是有一个有效的答案。
*/
func wiggleSort(nums []int)  {

	sort.Ints(nums)
	old:=nums
	var result =make([]int,0,len(nums))
	for len(nums) > 0 {
		if len(nums) >=2 {
			result = append(result, nums[0],nums[len(nums)-1])
			nums = nums[1:len(nums)-1]
		}else{
			result = append(result, nums[0])
			nums = nums[1:]
		}
	}

	copy(old,result)
}

