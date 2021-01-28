# 300. Longest Increasing Subsequence

## Problem

Given an unsorted array of integers, find the length of longest increasing subsequence.

**Example:**

```
Input: [10,9,2,5,3,7,101,18]
Output: 4 
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4. 
```

**Note:**

- There may be more than one LIS combination, it is only necessary for you to return the length.
- Your algorithm should run in O(*n2*) complexity.

**Follow up:** Could you improve it to O(*n* log *n*) time complexity?



## Solution

- 方法一：动态规划

  - 使用两层循环遍历数组：
    - 状态 `dp[i]` 表示的是以第 i 个数字结尾的最长递增子序列的长度
    - 第二层循环 j 表示的是最长递增子序列中的倒数第二个数，也就是 i 前面的那个数，则这个数的可能取值范围是从 0 到 i - 1，判断是否能够取值的条件是检查 `nums[i] > nums[j]` 是否满足，只有满足了 j 才是合法的。
    - 对于每一个合法的 j，通过检查 `dp[j]` 的值就可以知道以数字 j 结尾的最长递增子序列的长度，如果这个值大于 `dp[i]`，则就用 `dp[j] + 1` 来更新 `dp[i]`
    - 外层循环 i 每执行完一轮，就确定了一个以 i 结尾的最长递增子序列，则用该子序列的长度更新最终结果 res
  - 时间复杂度：$O(n^2)$，空间复杂度：$O(n)$

- 方法二：贪心 + 二分法优化时间复杂度：

  - 要使子字符串长度尽量长，则在构建结果的过程中，每次产生的子字符串我们都希望它的最后一个元素足够的小，这样之后的元素才更方便追加到后面；

  - 构造一个最长递增子序列 subq, 当我们要向这个子序列中插入一个元素时，找到在子序列中比当前数小的所有数里最大的那个数的位置 r ，然后用当前数替换掉 r + 1 位置上的值，这样就保证了插入的数是接在位置 r 后面的所有数中最小的那个

  - 时间复杂度：$O(nlogn)$， 空间复杂度：$O(n)$

    