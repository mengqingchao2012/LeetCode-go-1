# 659. Split Array into Consecutive Subsequences
## Problem

Given an array `nums` sorted in ascending order, return `true` if and only if you can split it into 1 or more subsequences such that each subsequence consists of consecutive integers and has length at least 3.

 

**Example 1:**

```
Input: [1,2,3,3,4,5]
Output: True
Explanation:
You can split them into two consecutive subsequences : 
1, 2, 3
3, 4, 5
```

**Example 2:**

```
Input: [1,2,3,3,4,4,5,5]
Output: True
Explanation:
You can split them into two consecutive subsequences : 
1, 2, 3, 4, 5
3, 4, 5
```

**Example 3:**

```
Input: [1,2,3,4,4,5]
Output: False
```

 

**Constraints:**

- `1 <= nums.length <= 10000`

## Solution

- 维护两个 map，一个用来统计各个数字出现的总次数（`countMap`），另一个用来保存分隔完毕后，以数字 i 结尾的数组共有多少个（`resMap`）
- 首先遍历一次数组，更新统计数字出现次数的 `countMap`
- 然后开始分组，遍历数组，对于元素 i ，检查 `resMap[i - 1]` 是否大于 0
  - 如果是，则将元素 i 加到 `resMap` 中 `i - 1` 数组的后面，即此时 `resMap[i - 1]--`， `resMap[i]++`
  - 优先添加到存在数组之后是因为题目要求分割后的数组长度必须要大于等于 3，如果对于任意新元素都选择新开一个数组的话，很容易产生非常多的短数组，最后反而得不到结果，因此选择将新加入的元素尽量添加到已经存在的数组中（贪心的思路）
  - 如果 `resMap[i - 1]` 不大于 0，则需要在 `resMap` 中新建数组，此时需要检查 `countMap[i + 1]` 和 `countMap[i + 2]` 是否大于 0，即一次性凑够 3 个数，更新 `resMap[i + 2]++`。如果凑不够的话就说明组成满足题意的结果，可以直接返回 false
- 时间复杂度：$O(n)$，空间复杂度：$O(n)$