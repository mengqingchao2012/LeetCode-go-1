# 260. 只出现一次的数字 III

## Problem

Given an array of numbers `nums`, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.

**Example:**

```
Input:  [1,2,1,3,2,5]
Output: [3,5]
```

**Note**:

1. The order of the result is not important. So in the above example, `[5, 3]` is also correct.
2. Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?

## Solution

- 思路：

  - 通过异或运算可以得到数组中只出现一次的两个数字异或后的结果
  - 异或运算的规则是：值相同的位异或结果为0，值不同的位异或结果为1，两个数字不同，则说明这两个数字中必然有至少一个位上的值是不相同的，则可以通过这个不同值的位来将数组分成两组，将这两个不同的数字打散到两个组中
  - 第一次遍历求得两个不同数字的异或值，然后使用掩码运算找到该值中第一个值为1的位
    - `x & (-x)` 求出的即是值 x 第一个值为1的位
  - 第二次遍历，分别用两个变量来存储两个不同组的异或结果，遍历到每个元素，都用该元素与上上一步求出来的掩码，用来将元素分到不同的组，最终的 x 和 y 的结果就是只出现一次的两个数字

- 时间复杂度：$O(n)$，空间复杂度：$O(1)$

  