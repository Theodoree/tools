package _00_399



/*

384. 打乱数组

打乱一个没有重复元素的数组。

示例:

// 以数字集合 1, 2 和 3 初始化数组。
int[] nums = {1,2,3};
Solution solution = new Solution(nums);

// 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
solution.shuffle();

// 重设数组到它的初始状态[1,2,3]。
solution.reset();

// 随机返回数组[1,2,3]打乱后的结果。
solution.shuffle();
*/
type Solution struct {
    cur   []int
}

func Constructor(nums []int) Solution {
    reset := make([]int, len(nums))
    copy(reset, nums)
    return Solution{
        cur:   nums,
    }

}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.cur
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {

    nums := make([]int, len(this.cur))
    copy(nums, this.cur)
    rand.Shuffle(len(nums), func(i int, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    })

    return nums
}
/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
