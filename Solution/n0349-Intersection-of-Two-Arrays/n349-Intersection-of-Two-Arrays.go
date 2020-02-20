package main

func intersection(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return []int{}
	}

	if n1 > n2 { //保证长度短的数组作为第一个参数，以减少后续遍历的元素个数，优化内存分配
		return intersection(nums2, nums1)
	}

	mp := make(map[int]int, n1)
	for _, v := range nums1 {
		if _, ok := mp[v]; !ok {
			mp[v] = 1
		}
	}

	res := make([]int, 0, n1)
	for _, v := range nums2 {
		if val, ok := mp[v]; ok && val != 0 {
			res = append(res, v)
			mp[v] = 0
		}
	}
	return res
}