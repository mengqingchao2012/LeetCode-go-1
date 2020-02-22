package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return []int{}
	}

	if n1 > n2 {
		return intersect(nums2, nums1)
	}

	mp := make(map[int]int, n1)
	for _, v := range nums1 {
		mp[v]++
	}

	res := make([]int, 0, n1)
	for _, v := range nums2 {
		if val, ok := mp[v]; ok && val != 0 {
			res = append(res, v)
			mp[v]--
		}
	}
	return res
}

//如果数组有序，则使用该方法：时间复杂度：O(mlogm + nlogn), 空间复杂度：O(1)
func intersect1(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return []int{}
	}

	sort.Sort(sort.IntSlice(nums1))
	sort.Sort(sort.IntSlice(nums2))

	i, j := 0, 0
	res := []int{}
	for i < n1 && j < n2 {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
