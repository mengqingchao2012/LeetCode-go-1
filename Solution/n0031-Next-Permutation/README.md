# 31. Next Permutation

## Problem

Implement **next permutation**, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be **[in-place](http://en.wikipedia.org/wiki/In-place_algorithm)** and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

```
`1,2,3` → `1,3,2`
`3,2,1` → `1,2,3`
`1,1,5` → `1,5,1`
```

## Solution

- 题意理解：

  - 给定一个数字，如果组成该数字的数字可以组合成一个更大的数字，则返回更大数字中最小的那个
  - 如果当前数字已经是最大数字，则返回其按照升序排序的数字

- 思路：

  - 要返回比原数字更大数字中最小的那个数字，则希望对数字的改动尽量的小，从右边开始往前找

  - 以数字 218421 为例，从右往前找，8421这个子数字是按照降序排列的，也就是说改变这个子数字已经无法得到一个更大的子数字了，而对于 18421 这个子数字，它不是按照降序排列的，因此可以找到一个比该数字要大的数字

  - 由此可知，需要调整的数字在 18421 的最高位 1 处，从最高位的右边数字里找到比该位大的所有数中最小的那个数，也就是2，交换两个数得到 28411

  - 为了得到最小的满足要求的数字，对于 28411，需要将 8411 变成值最小的组合方式，由于8411已经是按照降序排序的，所以可以从两边双指针向中间靠拢，对调元素，得到 1148

  - 由此，最终结果为：221148

  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$

    