# 986. Interval List Intersections
## Problem

Given two lists of **closed** intervals, each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

*(Formally, a closed interval `[a, b]` (with `a <= b`) denotes the set of real numbers `x` with `a <= x <= b`. The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented as a closed interval. For example, the intersection of [1, 3] and [2, 4] is [2, 3].)*

 

**Example 1:**

**![avatar](..\..\pic\l986.png)**

```
Input: A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
```

 

**Note:**

1. `0 <= A.length < 1000`
2. `0 <= B.length < 1000`
3. `0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9`

## Solution

- 思路：

  - 从题目可知：两个数组有序，且各自区间不相交

  - 合并两个数组的不同区间时，满足：

    ```
    start = Max(A[0], B[0])
    end = Min(A[1], B[1])
    ```

    且 start <= end

  - 因此使用双指针遍历两个数组，依次合并区间：

    - 得到一个交集区间后，较早结束的那个子区间不可能再与别的区间重叠，否则就与题目假设相矛盾，故向后移动其指针指向下一个子区间
    - 较长的区间还可能与别的区间产生交集，故要保持不动
    - 区间的结束可以通过 A[1] 和 B[1] 的大小来确定

- 时间复杂度：$O(m + n)$，空间复杂度：$O(m + n)$

- 其他：尽量不要用二分法求解区间问题，边界条件太难判断了