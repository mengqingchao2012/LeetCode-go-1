package main

// 保存前缀和，加快 SumRange 的执行速度
type NumArray struct {
	PrefixSum []int
}

// 为了统一边界，将前缀和定义为：s[i] 表示 0 ~ i-1 的元素的和
// 故：f(i, j) = s[j + 1] - s[i]
func Constructor(nums []int) NumArray {
	PrefixSum := make([]int, len(nums) + 1)
	PrefixSum[0] = 0
	for i := 1; i < len(PrefixSum); i++ { // 注意终止条件是遍历到 len(PrefixSum)
		PrefixSum[i] = PrefixSum[i - 1] + nums[i - 1]
	}
	return NumArray { PrefixSum }
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.PrefixSum[j + 1] - this.PrefixSum[i]
}
