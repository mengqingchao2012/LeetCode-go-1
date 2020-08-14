# 39. Combination Sum

## Problem

Given a **set** of candidate numbers (`candidates`) **(without duplicates)** and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sums to `target`.

The **same** repeated number may be chosen from `candidates` unlimited number of times.

**Note:**

- All numbers (including `target`) will be positive integers.
- The solution set must not contain duplicate combinations.

**Example 1:**

```
Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
```

**Example 2:**

```
Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
```



## Solution

思路：

- 方法一：思路来源于 sicp 的1.2.2中，换零钱问题：
  - 将总数为a的现金换成n种不同种类硬币的不同方式的总数为：
    - 将现金a换成除第一种硬币之外的所有其他硬币的不同的方式的数目，加上
    - 将现金a-d换成所有种类的硬币的不同方式数目，其中d是第一种硬币的币值；
  - 算法正确的理由：将换零钱分成两组：第一组全都没有使用第一种硬币，第二组都是用了第一种硬币。由此，换成零钱的全部方式的数目，等于完全不用第一种硬币的方式的数目，加上用了第一种硬币的换零钱方式的数目。后者的数目就等于去掉一个第一种硬币值后，剩下的现金数的换零钱方式的数目；
  - 递归终止条件：
    - 如果a的值为0，算作是找到了一种换零钱的方式；
    - 如果a小于0，说明没有找到换零钱的方式；
    - 如果n是0，说明也没有找到换零钱的方式；
- 方法二：回溯
  - 假设有 i 个数可以组成一个满足条件的结果，则这一组结果中的 i 位每一位都可以取数组中的任意值
  - 依次固定每一位的值，检查下一位的取值是否可以得到满足条件的结果
  - 如果可以，则将该组结果加入结果集，如果不行，则回退到上一位，再取下一个数来组成结果集
  - 在执行前对数组进行排序可以减少执行次数