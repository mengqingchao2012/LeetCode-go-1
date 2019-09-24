# LeetCode note

### 16. 最接近的三数之和

>给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
>
>例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
>
>与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
>

- 解法：

```go
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    res, min := 0, math.MaxInt64
    
    for i := len(nums) - 1; i >= 2; i-- {
        j, k := 0, i - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum == target {
                return sum
            } else if sum < target {
                j++
            } else {
                k--
            }
            
            diff := differ(sum, target)
            
            if diff < min {
                min = diff
                res = sum
            }
        }
    }
    return res
}

func differ(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}
```

- 思路：
  - 先将数组排序，然后从右边开始，固定住最后一个数，赋值为i，令j=0, k=i-1，即遍历i左边的区间，取出三个数求和并与target进行比较，若小于target，则增大j，若大于target则减小k，若相等则说明找到了要求的三个数；每轮循环结束都将i左移，直到最后区间内数字的个数小于3个，则结束；
  - 使用变量min来保存和与目标之间的差值的最小值，在循环中更新min；
  - 时间复杂度：$O(n^2)$，空间复杂度：$O(1)$