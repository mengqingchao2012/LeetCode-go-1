# 739. Daily Temperatures
## Problem

Given a list of daily temperatures `T`, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put `0` instead.

For example, given the list of temperatures `T = [73, 74, 75, 71, 69, 72, 76, 73]`, your output should be `[1, 1, 4, 2, 1, 1, 0, 0]`.

**Note:** The length of `temperatures` will be in the range `[1, 30000]`. Each temperature will be an integer in the range `[30, 100]`.

## Solution

- 解法一：辅助栈

  <img src="..\..\pic\l739-1.png" alt="avatar" style="zoom:25%;" />

  - 使用一个辅助栈，栈中存储的是元素的下标
  - 对于每一个元素，如果栈不为空，则将栈顶元素对应的值与当前元素的值进行比较，如果当前元素更大，则说明找到了第一个温度升高的日期，则用当前元素的下标减去栈顶元素存储的下标即是该元素的结果值。弹出栈顶元素继续比较
  - 如果栈中没有元素或者当前元素的值不大于栈顶元素对应的值，则将当前元素的下标入栈，继续遍历
  - 遍历完数组后，栈中遗留的元素都是找不到更大温度的元素，将这些元素的结果都设置为 0
  - 时间复杂度：$O(n)$（每个元素都入栈出栈了一次），空间复杂度：$O(n)$

- 解法二：跳跃遍历

  <img src="..\..\pic\l739-2.png" alt="avatar" style="zoom:35%;" />

  - 从右往左遍历，从倒数第二个元素开始，因为最后一个元素的结果肯定是 0

  - 每次遍历到一个新元素 i，首先比较 i 和 j = i + 1 的值，如果满足 T(j) > T(i)，则 res[i] = j - i，否则可以借助 res 中存储的结果，跳过那些比待比较元素（如 i + 1）小的元素，即更新 j += res[j]，然后再继续判断是否满足 T(j) > T(i)

  - 如果跳转到某个 res[j] = 0 的值时，说明找不到更大的元素了，可以直接设置 res[i] = 0

  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$

     

  