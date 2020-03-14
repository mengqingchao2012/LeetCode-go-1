# 300. Longest Increasing Subsequence

## Problems

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



## Solutions

- 方法一：动态规划

  <img src="..\..\pic\lc300.png" alt="avatar" style="zoom:50%;" />

  - 使用两层循环遍历数组：
    - 状态数组 d[i] 表示遍历到第 i 个元素时，最长子串的长度是多少；
    - 第二层循环 j 每次都从数组的第1个元素开始，一直遍历到外层循环所指的元素 i ，比较 nums[i] 和 nums[j] 的大小，如果 nums[i] > nums[j]， 则说明下标 i 对应的元素可以追加到下标 j 元素组成的结果集中，合成一个更大的子串，记录下这个子串的长度 cur 为 d[j] + 1，然后从 cur 和 d[i] 中取较大的那一个，就表示遍历到第 i 个元素时，组成的最长子串的长度
    - 每次得到一个确定的 d[i] 后，要从数组 d 中取出最大的那个值作为最终的结果值
  - 时间复杂度：$O(n^2)$，空间复杂度：$O(n)$

- 方法二：动态规划 + 二分法优化时间复杂度：

  - 要使子字符串长度尽量长，则在构建结果的过程中，每次产生的子字符串我们都希望它的最后一个元素足够的小，这样之后的元素才更方便追加到后面；

  - 状态数组 d 表示最终的结果集；

  - 遍历数组，使用二分法将遍历到的元素插入到结果集数组 d 中，最终数组 d 的长度就是最长子序列的长度

  - 时间复杂度：$O(nlogn)$， 空间复杂度：$O(n)$

    