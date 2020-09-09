# 75. Sort Colors
## Problem

Given an array with *n* objects colored red, white or blue, sort them **[in-place](https://en.wikipedia.org/wiki/In-place_algorithm)** so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

**Note:** You are not suppose to use the library's sort function for this problem.

**Example:**

```
Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

**Follow up:**

- A rather straight forward solution is a two-pass algorithm using counting sort.
  First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
- Could you come up with a one-pass algorithm using only constant space?

## Solution

- 荷兰国旗问题
- 注意到，题目中给出的颜色是有限的，即只有 0， 1， 2 三种，原地排序后，最前面的肯定是 0，最后面的肯定是 2，1 夹在中间
- 使用双指针分别指向数组的两端，然后定义一个游标 i，从数组开头开始遍历：
  - 如果 nums[i] == 0，则交换 nums[left] 和 nums[i] 的值，并同时将 i 和 left 指针后移
  - 如果 nums[i] == 1，则只需要将 i 指针后移，因为任意元素 1 的位置就是在中间
  - 否则交换 nums[right] 和 nums[i] 的值，注意此时只能左移 right 指针，因为交换后 nums[i] 的值是不确定的，还需要下一轮继续比较
- 时间复杂度：$O(n)$，空间复杂度：$O(1)$