package main

import (
	"LeetCode-go/utils"
	"math"
)

//方法一
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	midNum := (n1 + n2 + 1) >> 1

	left, right := 0, n1
	for left < right {
		partA := left + ((right - left) >> 1)
		partB := midNum - partA
		if nums1[partA] < nums2[partB - 1] {
			left = partA + 1
		} else {
			right = partA
		}
	}

	m1 := left
	m2 := midNum - left

	var v1, v2 int
	if m1 <= 0 {
		v1 = math.MinInt32
	} else {
		v1 = nums1[m1 - 1]
	}
	if m2 <= 0 {
		v2 = math.MinInt32
	} else {
		v2 = nums2[m2 - 1]
	}

	c1 := utils.Max(v1, v2)
	if (n1 + n2) % 2 == 1 {
		return float64(c1)
	}

	if m1 >= n1 {
		v1 = math.MaxInt32
	} else {
		v1 = nums1[m1]
	}
	if m2 >= n2 {
		v2 = math.MaxInt32
	} else {
		v2 = nums2[m2]
	}

	c2 := utils.Min(v1, v2)
	return float64(c1 + c2) * 0.5
}

// 方法二
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if (total & 1) == 1 {
		return findKthSmallestInSortedArrays(&nums1, &nums2, total / 2 + 1)
	} else {
		c1 := findKthSmallestInSortedArrays(&nums1, &nums2, total / 2 + 1)
		c2 := findKthSmallestInSortedArrays(&nums1, &nums2, total / 2)
		return (c1 + c2) * 0.5
	}
}

func findKthSmallestInSortedArrays(nums1, nums2 *[]int, k int) float64 {
	n1, n2 := len(*nums1), len(*nums2)
	delt1, delt2 := 0, 0

	for {
		if n1 == 0 {
			return float64((*nums2)[delt2 + k - 1])
		}
		if n2 == 0 {
			return float64((*nums1)[delt1 + k - 1])
		}
		if k == 1 {
			return float64(utils.Min((*nums1)[delt1], (*nums2)[delt2]))
		}

		i := utils.Min(k / 2, n1)
		j := utils.Min(k - i, n2)
		a, b := (*nums1)[delt1 + i - 1], (*nums2)[delt2 + j - 1]

		if i + j == k && a == b {
			return float64(a)
		}

		if a <= b {
			delt1 += i
			n1 -= i
			k -= i
		}

		if a >= b {
			delt2 += j
			n2 -= j
			k -= j
		}
	}
}

//func main() {
//	a := []int{1, 2, 5}
//	b := []int{3, 4}
//
//	res := findMedianSortedArrays(a, b)
//	fmt.Println(res)
//}
