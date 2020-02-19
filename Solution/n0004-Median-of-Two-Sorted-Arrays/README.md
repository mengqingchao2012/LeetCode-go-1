# 4. Median of Two Sorted Arrays

## Problem
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

## Solution
- 解法一：中位数其实就是求有序数组中中间的那个，可以看做是求第k小问题，需要注意的是，数组长度为奇数是，中位数可以直接得到；
    如果数组长度是偶数，中位数等于中间两个数的平均值
    
    -   思路：实现一个算法求第k小，然后就可以分奇数和偶数情况分别求中位数
    
    -   难点在于数据是分散在两个数组中（具体解法看代码注释）
    
    -   时间复杂度：O(log(m+n))，空间复杂度：O(1)