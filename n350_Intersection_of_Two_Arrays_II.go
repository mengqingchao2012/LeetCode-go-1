package main

import "sort"

//方法一：使用排序法：Time：O(mlogm + nlogn)，Space：O(1)
func intersect(nums1 []int, nums2 []int) []int {
	//首先对两个数组排序
	//排序也可以使用sort.Slice(nums, func(i, j int) bool {}) 来实现
	sort.Sort(sort.IntSlice(nums1))
	sort.Sort(sort.IntSlice(nums2))

	i, j := 0, 0
	var res []int
	//使用i,j两个下标分别遍历两个数组，因为数组有序，如果i指向的值小于j，则需要将i后移，同理处理j
	for i < len(nums1) && j < len(nums2) {
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

//方法二：使用哈希表：Time：O(m + n)，Space：O(m)
func intersect1(nums1 []int, nums2 []int) []int {

	count := make(map[int]int)

	for idx := range nums1 { //遍历nums1，统计各元素出现的次数
		count[nums1[idx]]++
	}

	var res []int
	for _, k := range nums2 { //遍历nums2，检查元素是否在哈希表中，如果在，就将元素加入结果集，哈希表中计数减一
		if val, ok := count[k]; ok && val != 0 {
			res = append(res, k)
			count[k]--
		}
	}
	return res
}


