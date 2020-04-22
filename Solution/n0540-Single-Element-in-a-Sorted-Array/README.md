# 540. Single Element in a Sorted Array

## Problem

You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once. Find this single element that appears only once.

 
**Example 1:**

```
Input: [1,1,2,3,3,4,4,8,8]
Output: 2
```

**Example 2:**

```
Input: [3,3,7,7,10,11,11]
Output: 10
```

 
**Note:** Your solution should run in O(log n) time and O(1) space.


## Solution

- 使用二分法：

  - 利用数组的有序性可知：相同的元素必定相邻，即当前元素如果不是单身元素，其要么等于前一个元素，要么等于后一个元素，如果都不相等，则找到了单身元素
  - 使用二分法求出中间元素后：
    - 如果其等于前一个元素，则说明当前元素是重复元素的第二个元素，将mid移到第一个元素上，即mid-1
    - 如果其等于后一个元素，则说明当前元素是重复元素的第一个元素，保持mid值不变
    - 此时通过 mid - low 可以计算出前面的元素个数，如果是奇数，则说明单身元素在前半部分元素里，更新 high = mid - 1
    - 否则说明单身元素在后半部分元素里，更新 low = mid + 2，注意是 +2，因为要跳过重复的元素

- 时间复杂度：$O(logn)$，空间复杂度：$O(1)$

  