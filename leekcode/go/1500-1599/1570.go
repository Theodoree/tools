package _500_1599

/*
1570. 两个稀疏向量的点积
给定两个稀疏向量，计算它们的点积（数量积）。

实现类 SparseVector：

SparseVector(nums) 以向量 nums 初始化对象。
dotProduct(vec) 计算此向量与 vec 的点积。
稀疏向量 是指绝大多数分量为 0 的向量。你需要 高效 地存储这个向量，并计算两个稀疏向量的点积。

进阶：当其中只有一个向量是稀疏向量时，你该如何解决此问题？
*/

type SparseVector struct {
	buf []int
}

func Constructor(nums []int) SparseVector {
	var m SparseVector
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 高位数字,低位索引,该题目,nums[i]最大为100,长度最为10^5
			m.buf = append(m.buf, i|(nums[i]<<32))
		}
	}

	return m
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	var sum int

	left := this.buf
	right := vec.buf
	for len(left) > 0 && len(right) > 0 {
		lIdx := left[0] & math.MaxInt32
		rIdx := right[0] & math.MaxInt32
		if lIdx < rIdx {
			left = left[1:]
		} else if rIdx < lIdx {
			right = right[1:]
		} else {
			sum += left[0] >> 32 * right[0] >> 32
			left = left[1:]
			right = right[1:]
		}
	}
	return sum
}
