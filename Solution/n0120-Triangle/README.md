# 120. Triangle
## Problem

Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

```
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
```


The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:

Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

## Solution

- 方法一：自顶向下

  - 设状态 `dp[i][j]` 表示到达 `(i, j)` 节点是的最小路径和
  - `dp[i][j] = Min(dp[i - 1][j - 1], dp[i - 1][j]) + a[i][j]`
  - 特殊情况：
    - 对于三角形的左边界：`dp[i][0] = dp[i - 1][0] + a[i][0]`
    - 对于三角形的右边界：`dp[i][i] = dp[i - 1][i - 1] + a[i][i]`
    - 注意：三角形中每一行的元素个数等于行数，最后一个元素行列坐标相同
  - 求出所有状态后，遍历 `dp` 的最后一行，找到最小值即可
  - 时间复杂度：$O(m^2)$，空间复杂度：$O(m^2)$，m 表示三角形的行数

- 方法二：自底向上

  - 与自顶向下刚好相反，自底向上求解时，有：

    `dp[i][j] = Min(dp[i + 1][j], dp[i + 1][j + 1]) + a[i][j]`

  - 特殊情况只有最后一行，其值就等于原数组中最后一行的值

  - 自底向上的方法对边界条件的兼容性更好，每一行都是从 0 遍历到 i

  - 最终返回值只有一个，就是最小值

  - 时间复杂度：$O(m^2)$，空间复杂度：$O(m^2)$，m 表示三角形的行数

- 方法三：在方法二的基础上，可以将 `dp` 的第一维删掉，将空间复杂度降为 $O(n)$

  