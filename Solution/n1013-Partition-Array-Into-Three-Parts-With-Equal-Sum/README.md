# 1013. Partition Array Into Three Parts With Equal Sum

## Problem

Given an array `A` of integers, return `true` if and only if we can partition the array into three **non-empty** parts with equal sums.

Formally, we can partition the array if we can find indexes `i+1 < j` with `(A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1])`

 
**Example 1:**

```
Input: A = [0,2,1,-6,6,-7,9,1,2,0,1]
Output: true
Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
```

**Example 2:**

```
Input: A = [0,2,1,-6,6,7,9,-1,2,0,1]
Output: false
```

**Example 3:**

```
Input: A = [3,3,6,5,-2,2,5,1,-9,4]
Output: true
Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
```

 
**Constraints:**

- `3 <= A.length <= 50000`
- `-10^4 <= A[i] <= 10^4`


## Solution

- 思路：

  - 先遍历一遍数组求和，然后将和除以3，如果不能整除，直接返回 false，能整除，则除得的结果就是每部分的和
  - 双指针两头遍历，分别对左右两边求和，如果在左右两边都求得部分和时，两个指针还没有重叠，则说明有解，返回 true
  - 要注意，因为必须分成非空的三部分，所以满足部分和条件时，仅仅两个指针不重叠是不能满足需求的，还要加上一个不相邻的条件，即保证中间部分不为空，具体实现时，将循环退出条件改为 left + 1 < right 即可
  - 坑点：
    - 要将左右两边和初始化为两端元素，在中间过程中累加时，要先移动下标，再取值相加，否则会出错

- 时间复杂度：$O(n)$，空间复杂度：$O(1)$

  