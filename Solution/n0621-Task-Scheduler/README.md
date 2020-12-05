# 621. Task Scheduler
## Problem

Given a characters array `tasks`, representing the tasks a CPU needs to do, where each letter represents a different task. Tasks could be done in any order. Each task is done in one unit of time. For each unit of time, the CPU could complete either one task or just be idle.

However, there is a non-negative integer `n` that represents the cooldown period between two **same tasks** (the same letter in the array), that is that there must be at least `n` units of time between any two same tasks.

Return *the least number of units of times that the CPU will take to finish all the given tasks*.

 

**Example 1:**

```
Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation: 
A -> B -> idle -> A -> B -> idle -> A -> B
There is at least 2 units of time between any two same tasks.
```

**Example 2:**

```
Input: tasks = ["A","A","A","B","B","B"], n = 0
Output: 6
Explanation: On this case any permutation of size 6 would work since n = 0.
["A","A","A","B","B","B"]
["A","B","A","B","A","B"]
["B","B","B","A","A","A"]
...
And so on.
```

**Example 3:**

```
Input: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
Output: 16
Explanation: 
One possible solution is
A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> idle -> idle -> A -> idle -> idle -> A
```

 

**Constraints:**

- `1 <= task.length <= 104`
- `tasks[i]` is upper-case English letter.
- The integer `n` is in the range `[0, 100]`.

## Solution

- 思路：

  - 以 `tasks=[A, A, A, B, B, B, D]，n=2` 为例

  - 使用出现频率最高的重复任务来进行分组，此例中 A=B=3，选择 A，相同任务的冷却时间为 n，则在每组中，使用其他任务进行填充，如果没有更多的任务，则填入 idle。每个分组的执行时间为 n + 1

  - 如果在填充完所有由频率最高的元素组成的分组后，还有别的单个任务怎么办，如：`tasks=[A, A, A, B, B, B, D, X, Y, Z, X, Y, Z]，n=2`，此时实际上还是可以基于最大频率创建的分组，在每组中间穿插填入后面多出来的元素即可，即此时每个分组的执行时间可能大于 n + 1

    <img src="..\..\pic\l621.png" alt="avatar" style="zoom:50%;" />

- 方法一：最大堆

  - 维护一个最大堆，对每个元素求出其频率后，加入堆中
  - 使用 res, idle 两个变量，当堆中元素不为空时，取出 n + 1 个元素，分别将其频率减 1，如果频率不为 0，则继续加入堆中，更新 `res += n + 1`，`idle = n + 1 - 当前堆中元素`，注意 idle 的作用只是用于确定最后一次任务到底消耗了多少时间片，中间并没有其他用处，因此即便中间计算过程中会出现负值也不影响
  - 当堆中元素为空时，则说明完成了最终计算，最后的结果等于 `res - idle`
  - 时间复杂度：
    - 虽然存在堆中的是频率，但实际上每次弹出堆顶元素，频率减一后又放回，到最后堆中元素为空，总共执行了 m 次，m 是 tasks 数组的长度，每次弹出和放回的时间复杂度都是 log26，因为堆中最多只有 26 个元素，因此时间复杂度为 $O(mlog26) = O(m)$
  - 空间复杂度是 O(1)，堆中最多只有 26 个元素，复制哈希表也最多只有 26 个元素，因此是常量
  
- 方法二：数学计算

  - 由之前的分析可知，通过频率最大的元素个数和间隔时间可以将数组进行分组，每组需要消耗的时间片为 n + 1，故除了最后一组之外，前面几组的开销为 `(maxFreq - 1) * (n + 1)`，对于最后一组元素，实际上其开销等于出现频率等于最大频率的元素个数，可以通过一个额外的变量 `cnt` 来统计
    - 如 `AAABBBD` 这个任务中，最后一组是 `AB`，因为 A 和 B 的频率都是 3，而 D 只有 1
  - 因此最终占用的时间片为 `(maxFreq - 1) * (n + 1) + cnt`
  - 如果出现添加了 `XYZXYZ` 后的情况，即元素总个数大于 `maxFreq * (n + 1)`，此时则总的时间片开销即为数组长度
  - 时间复杂度为：$O(n)$，空间复杂度为：$O(1)$

