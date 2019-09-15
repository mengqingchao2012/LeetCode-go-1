package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for item, n := range nums {
		if v, ok := mp[target - n]; ok {
			return []int{v, item}
		}
		mp[n] = item
	}
	return nil
}