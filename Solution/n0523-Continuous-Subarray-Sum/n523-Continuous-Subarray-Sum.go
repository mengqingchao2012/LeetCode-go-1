package main

func checkSubarraySum(nums []int, k int) bool {
	length := len(nums)
	if length == 0 { return false }

	prefixSum := map[int]int{}
	prefixSum[0] = -1
	sum, mod := 0, 0

	for i := 0; i < length; i++ {
		sum += nums[i]
		if k == 0 {
			mod = sum
		} else {
			mod = sum % k
		}

		if v, ok := prefixSum[mod]; ok {
			if i - v > 1 {
				return true
			}
		} else {
			prefixSum[mod] = i
		}
	}
	return false
}
