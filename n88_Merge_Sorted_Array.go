package main

//Time：O(n)，Space：O(1)
func merge(nums1 []int, m int, nums2 []int, n int)  {
	k := m + n - 1
	i, j := m - 1, n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}
