# 41. First Missing Positive
## Problem

Given an unsorted integer array, find the smallest missing positive integer.

**Example 1:**

```
Input: [1,2,0]
Output: 3
```

**Example 2:**

```
Input: [3,4,-1,1]
Output: 2
```

**Example 3:**

```
Input: [7,8,9,11,12]
Output: 1
```

**Follow up:**

Your algorithm should run in *O*(*n*) time and uses constant extra space.

## Solution

- Cyclist Sort：
  - 要注意几个点：
    - 数组中的数字是任意加入的，不一定从 0 或者 1 开始
    - 数组中元素的值可能大于等于数组的有效下标
    - 正数从 1 开始