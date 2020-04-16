package main

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	mp := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := mp[v]; ok {
			return false
		}
		mp[v] = struct{}{}
	}
	return true
}
