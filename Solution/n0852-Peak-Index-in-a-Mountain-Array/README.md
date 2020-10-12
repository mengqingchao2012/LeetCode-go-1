# 852. Peak Index in a Mountain Array
## Problem

Let's call an array `arr` a **mountain** if the following properties hold:

- `arr.length >= 3`
- There exists some `i` with `0 < i < arr.length - 1` such that:
  - `arr[0] < arr[1] < ... arr[i-1] < arr[i]`
  - `arr[i] > arr[i+1] > ... > arr[arr.length - 1]`

Given an integer array arr that is **guaranteed** to be a mountain, return any `i` such that `arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`.

 

**Example 1:**

```
Input: arr = [0,1,0]
Output: 1
```

**Example 2:**

```
Input: arr = [0,2,1,0]
Output: 1
```

**Example 3:**

```
Input: arr = [0,10,5,2]
Output: 1
```

**Example 4:**

```
Input: arr = [3,4,5,1]
Output: 2
```

**Example 5:**

```
Input: arr = [24,69,100,99,79,78,67,36,26,19]
Output: 2
```

 

**Constraints:**

- `3 <= arr.length <= 104`
- `0 <= arr[i] <= 106`
- `arr` is **guaranteed** to be a mountain array.

## Solution

- 思路：
  - 数组是局部有序的，即：前半部分升序，后半部分降序
  - 使用二分法找到中间结点 mid，通过比较 arr[mid] 和 arr[mid + 1] 的值即可知晓当前 mid 所在的位置
    - arr[mid] > arr[mid + 1]：说明 mid 在后半部分，则最大值要么是 mid，要么在 mid 之前，因此更新 high = mid
    - 反之，则说明 mid 在前半部分，最大值必然在 mid 之后，更新 low = mid + 1
  - 当 low == high 时，说明收敛到了最大值，返回