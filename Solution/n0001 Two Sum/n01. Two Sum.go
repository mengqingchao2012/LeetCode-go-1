package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for idx, value := range nums {
		if v, ok := mp[target - value]; ok {
			return []int{v, idx}
		} else {
			mp[value] = idx
		}
	}
	return nil
}


