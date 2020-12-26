# 84. Largest Rectangle in Histogram
## Problem

Given *n* non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.

 

![avatar](..\..\pic\l84-1.png)
Above is a histogram where width of each bar is 1, given height = `[2,1,5,6,2,3]`.

 

![avatar](..\..\pic\l84-2.png)
The largest rectangle is shown in the shaded area, which has area = `10` unit.

 

**Example:**

```
Input: [2,1,5,6,2,3]
Output: 10
```

## Solution

- 对于某个高度，使用单调栈分别求出其左边界和右边界，则该高度下，组成矩形的面积就可以求出来了
- 求边界实际上就是求在某个柱子的左/右边，第一个高度小于该柱子高度的下标
- 时间复杂度：$O(n)$，空间复杂度：$O(n)$