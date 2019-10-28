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



### 17.电话号码的数字组合

>给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
>
>给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
>
>
>
>示例:
>
>输入："23"
>输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
>说明:
>尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

- 解法一：递归法：

```go
var mapping = map[byte]string {
	0: "abc",
	1: "def",
	2: "ghi",
	3: "jkl",
	4: "mno",
	5: "pqrs",
	6: "tuv",
	7: "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	var result []string
	recursion(digits, 0, "", &result)
	return result

}
//辅助函数，用于递归求解
func recursion(digits string, idx int, str string, result *[]string) { //注意数组要传入指针，因为最后要取到返回值
	if idx == len(digits) { //递归的退出条件
		*result = append(*result, str) //注意此处的解引用符号不能漏
		return
	}

	chars := mapping[digits[idx] - '2']
	for i := 0; i < len(chars); i++ {
		nxtStr := str + string(chars[i])
		recursion(digits, idx + 1, nxtStr, result)
	}
}
```

- 思路：
  - 递归解法一般是DFS的，迭代解法一般是BFS的；
  - 对于每个位置，穷举其所有可能得结果，然后回溯到下一个位置继续穷举，比如：`23`这个数字，`2` 的可能取值为 `abc`，当取`a`时，下一位 `3` 的可能取值为 `def`，即产生三个解`ad, ae, af`，回溯，然后`2`取值为`b`，继续穷举；
  - 时间复杂度：$O(4^n)$
    - 最坏情况：数字为`79`时，此时每一位上都有四种可能性（pqrs * wxyz），第一个位置要调用4次，第二个位置要调用 $4^2$ 次，第n个位置要调用 $4^n$ 次，则总时间为：$4 + 4^1 + ... + 4^n$ 
  - 空间复杂度：$O(n)$：取决于最大递归深度，n个字母组合时，最大字母长度为n，递归深度也为n
- 方法二：迭代法：

```go
func letterCombinations1(digits string) []string {
	if digits == "" {
		return nil
	}
	res := []string{""} //注意初始化，此时res的len为1

	for i := 0; i < len(digits); i++ {
		chars := mapping[digits[i] - '2']
		var temp []string
        //对于res中的每一个值，都将其与下一个值的所有可能值进行组合
		for j := 0; j < len(res); j++ { 
			for k := 0; k < len(chars); k++ {
				temp = append(temp, res[j] + string(chars[k]))
			}
		}
		res = temp
	}

	return res
}
```



### 31.下一个排列

>实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
>
>如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
>
>必须原地修改，只允许使用额外常数空间。
>
>以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
>1,2,3 → 1,3,2
>3,2,1 → 1,2,3
>1,1,5 → 1,5,1
>

- 解法：

```go
func nextPermutation(nums []int)  {
    if nums == nil {
        return 
    }
    
    n := len(nums)
    p := n - 2
    
    for p >= 0 && nums[p] >= nums[p + 1] {
        p--
    }
    
    if p >= 0 {
        i := n - 1
        for i > p && nums[i] <= nums[p] {
            i--
        }
        nums[p], nums[i] = nums[i], nums[p]
    }
    
	j := p + 1
	for i := n - 1; j < i; {
		nums[j], nums[i] = nums[i], nums[j]
		j++
		i--
	}
}
```

- 思路：
  - 由题意可得：要找的数是比当前数大的所有数中最小的那个；
  - 因为要找最小的数，所以从低位开始找：设置指针p指向倒数第二位，依次比较当前指针所指的数和他的低一位数的大小，并前移指针，当找到第一个不满足`nums[p] >= nums[p + 1]`的数时，指针P所指的数即是需要交换位置的数。此时再从低位开始遍历数组，找到第一个比`nums[p]`大的数`nums[i]`，并将两个数交换，最后再将p + 1和i之间的所有数首尾互换（即降序排列）即可；
  - 如数组为：`[2, 1, 8, 4, 2, 1]`，找到指针p的位置为8前面的1，找到第一个比1大的数为2，交换1和2以后，得到`[2, 2, 8, 4, 1,1]`，此时再将数字8到数组结束之间的所有数首尾两两互换，最终可得：`[2, 2, 1, 1, 4, 8]`
  - 时间复杂度：$O(n)$， 空间复杂度：$O(1)$



### 34. 在排序数组中查找元素的第一个和最后一个位置

>给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
>
>你的算法时间复杂度必须是 O(log n) 级别。
>
>如果数组中不存在目标值，返回 [-1, -1]。
>
>示例 1:
>
>输入: nums = [5,7,7,8,8,10], target = 8
>输出: [3,4]
>示例 2:
>
>输入: nums = [5,7,7,8,8,10], target = 6
>输出: [-1,-1]

- 解法：

```go
func searchRange(nums []int, target int) []int {
    size := len(nums)
    if size == 0 {
        return []int{-1, -1}
    }
    
    end := findLastIndex(&nums, target)
    begin := findLastIndex(&nums, target - 1) + 1
    
    if begin >= 0 && begin <=end && end < size {
        return []int{begin, end}
    }
    
    return []int{-1, -1}
}

func findLastIndex(nums *[]int, target int) int {
    size := len(*nums)
    low, high := 0, size - 1
    
    for low <= high {
        mid := low + (high - low >> 1)
        if (*nums)[mid] > target {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return high
}
```

- 思路：

  - 二分法的变体；

  - 结束位置的找法：二分法找到等于目标值的一个下标后，继续向右查找，即，high不变，low每次加1，直到循环条件不再满足退出，此时的high的值即是要查找的目标的结束位置；

    - 实现也很简单，只需修改循环更新条件，只分 `nums[mid] > target` 和 小于等于 target 的情况来写即可；

  - 开始位置的找法：找值比目标值小1的数的结束位置，再加1即可；

    - 原理：即便比目标值小1的数在数组中不存在，使用`findLastIndex`函数在目标值不存在时，返回的 high 的值是数组中小于目标值的所有数中最大的那一个，因此在得到的这个下标上加1即可得到目标值的开始位置；

  - 时间复杂度：$O(logn)$，空间复杂度：$O(1)$

    

### 39. 组合总和

>给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
>
>candidates 中的数字可以无限制重复被选取。
>
>说明：
>
>所有数字（包括 target）都是正整数。
>解集不能包含重复的组合。 
>示例 1:
>
>输入: candidates = [2,3,6,7], target = 7,
>所求解集为:
>[
>  [7],
>  [2,2,3]
>]
>示例 2:
>
>输入: candidates = [2,3,5], target = 8,
>所求解集为:
>[
>  [2,2,2,2],
>  [2,3,3],
>  [3,5]
>]
>

- 解法：

```go
import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var solution []int

	combsum(candidates, solution, target, &result)
	return result
}

func combsum(candidates, solution []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, solution)
		return
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

    //注意这里重新再切切片是为了保证下一次append的时候，solution产生一个新的底层数组，如果不执行这一步的话，由于切片共享底层数组会导致最终得出错误的答案
	solution = solution[: len(solution) : len(solution)]

	combsum(candidates, append(solution, candidates[0]), target - candidates[0], result)
	combsum(candidates[1:], solution, target, result)
}
```

- 思路：
  - 思路来源sicp的1.2.2中，换零钱问题：
  - 将总数为a的现金换成n种不同种类硬币的不同方式的总数为：
    - 将现金a换成除第一种硬币之外的所有其他硬币的不同的方式的数目，加上
    - 将现金a-d换成所有种类的硬币的不同方式数目，其中d是第一种硬币的币值；
  - 算法正确的理由：将换零钱分成两组：第一组全都没有使用第一种硬币，第二组都是用了第一种硬币。由此，换成零钱的全部方式的数目，等于完全不用第一种硬币的方式的数目，加上用了第一种硬币的换零钱方式的数目。后者的数目就等于去掉一个第一种硬币值后，剩下的现金数的换零钱方式的数目；
  - 递归终止条件：
    - 如果a的值为0，算作是找到了一种换零钱的方式；
    - 如果a小于0，说明没有找到换零钱的方式；
    - 如果n是0，说明也没有找到换零钱的方式；



### 47.全排列II

>给定一个可包含重复数字的序列，返回所有不重复的全排列。
>
>示例:
>
>输入: [1,1,2]
>输出:
>[
>  [1,1,2],
>  [1,2,1],
>  [2,1,1]
>]
>

- 解法：

```go
import "sort"

func permuteUnique(nums []int) [][]int {
	if nums == nil {
		return [][]int{}
	}

	var res [][]int

	sort.Ints(nums)
	res = append(res, nums)
	for isFind, num := nextPermutations(nums); isFind == true;{
		res = append(res, num)
		isFind, num = nextPermutations(num)
	}
	return res
}

//注意不要传切片指针进来，切片指针共享底层数组，在辅助函数中改写切片指针会导致所有结果都受到影响
func nextPermutations(num []int) (bool, []int) {
	n := len(num)
	p := n - 2

	nums := make([]int, n)
	copy(nums, num)

	for p >= 0 && nums[p] >= nums[p + 1] {
		p--
	}

	if p >= 0 {
		i := n - 1
		for i > p && nums[i] <= nums[p] {
			i--
		}
		nums[p], nums[i] = nums[i], nums[p]
	}

	j := p + 1
	for i := n - 1; j < i; {
		nums[j], nums[i] = nums[i], nums[j]
		j++
		i--
	}

	return p != -1, nums
}
```

- 思路：
  - 基于第31题，下一个排列的解法求解；
  - 对数组按升序排序，然后把数组转换成字典序中下一个更大的排列，并把生成的结果加入结果集中，这样一来就完全避免了重复结果的出现；
  - 如：`[1, 1, 2]`这个数组生成的最终结果为：`[[1, 1, 2], [1, 2, 1], [2, 1, 1]]`;
  - 时间复杂度：$O(n * n!)$，空间复杂度：$O(1)$



### 61. 旋转链表

>给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
>
>示例 1:
>
>输入: 1->2->3->4->5->NULL, k = 2
>输出: 4->5->1->2->3->NULL
>解释:
>向右旋转 1 步: 5->1->2->3->4->NULL
>向右旋转 2 步: 4->5->1->2->3->NULL
>示例 2:
>
>输入: 0->1->2->NULL, k = 4
>输出: 2->0->1->NULL
>解释:
>向右旋转 1 步: 2->0->1->NULL
>向右旋转 2 步: 1->2->0->NULL
>向右旋转 3 步: 0->1->2->NULL
>向右旋转 4 步: 2->0->1->NULL
>

- 解法：构建成环

```go
type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }
    
    oldTail := head
    var n = 1 //注意n从1开始计数
    for ; oldTail.Next != nil; n++ {
        oldTail = oldTail.Next
    }
    oldTail.Next = head //头尾相连，构建成环
    
    newTail := head
    for i := 0; i < n - k % n - 1; i++ {
        newTail = newTail.Next
    }
    res := newTail.Next
    newTail.Next = nil
    
    return res
}
```

- 思路：
  - 构建成环，则旋转操作其实就变为找到新的头结点和尾结点，然后断开链表；
  - 新链表头的位置：`n - k % n`，链表尾的位置：`n - k % n - 1`，n是链表中结点的个数，注意n从1开始数才对；
  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$



### 66. 加一

>给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
>
>最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
>
>你可以假设除了整数 0 之外，这个整数不会以零开头。
>
>示例 1:
>
>输入: [1,2,3]
>输出: [1,2,4]
>解释: 输入数组表示数字 123。
>示例 2:
>
>输入: [4,3,2,1]
>输出: [4,3,2,2]
>解释: 输入数组表示数字 4321。
>

- 解法：

```go
func plusOne(digits []int) []int {
    carry := 1
    for i := len(digits) - 1; i >= 0; i-- {
        sum := digits[i] + carry
        digits[i] = sum % 10
        carry = sum / 10
    }
    if carry != 0 {
        return append([]int{1}, digits...)
    }
    return digits
}
```

- 思路：
  - 关键点：进位会导致相关位上的数值变化，数组的长度也可能会发生变化；
  - 将求和后的结果分成两部分，保留在原位上的数值和要进位的数值；
  - 初始化时，将加一操作抽象为更低位的进位值，即令carry=1，逆序遍历数组，将数组中的值与carry值相加，和对10取余即得到当前位上的数值，和对10取商，即得到进位的数值；
  - 时间复杂度为：$O(n)$
  
    

### 287.寻找重复数

>给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
>
>示例 1:
>
>输入: [1,3,4,2,2]
>输出: 2
>示例 2:
>
>输入: [3,1,3,4,2]
>输出: 3
>说明：
>
>不能更改原数组（假设数组是只读的）。
>只能使用额外的 O(1) 的空间。
>时间复杂度小于 O(n2) 。
>数组中只有一个重复的数字，但它可能不止重复出现一次。
>

- 方法一：二分法

```go
func findDuplicate(nums []int) int {
	low, high := 1, len(nums) - 1 //注意low从1开始
	for low < high { //注意边界
		mid := low + ((high - low) >> 1)
		count := 0
		for _, num := range nums {
			if num <= mid { //注意边界
				count++
			}
		}

		if count > mid { //注意边界
			high = mid
		} else {
			low = mid + 1 //注意边界
		}
	}

	return low
}
```

- 思路：
  - 题目理解：数组的长度：n+1，数字范围：[1, n]，由此可知，数组中是一定存在重复数字的，比如，n取5，则数组长度为6个数，数字范围：[1, 5]只有5个取值；
  - 用到的数学方法：抽屉原理，简单理解：一个萝卜一个坑，找到中间那个，分别统计左右两边数的个数，如果不存在重复数字，则左右两边个数相同，等于中间数，否则说明多的那一边存在重复数字；
  - low和high不是下标，而是取值范围中的两个极值，mid求出的即是中间数；
  - 统计数组中值小于等于中间数值的数字个数count，如果count的个数大于中间数值，说明重复数值在小于中间值的这一边，反之亦然；
  - 迭代更新low和high，直到两个数相等，此时这个值就是重复数；
  - 时间复杂度：$O(nlog(n))$，空间复杂度：$O(1)$；
- 方法二：模拟带环链表：

```go
func findDuplicate(nums []int) int {
    slow := nums[0]
    fast := nums[nums[0]]
    
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    
    ptr := 0
    for slow != ptr {
        slow = nums[slow]
        ptr = nums[ptr]
    }
    
    return slow
}
```

- 思路：
  - 由于数值的取值范围小于数组的大小，由此，数值的取值可以当做数组的下标不会越界；
  - 每次以当前的值作为下标，取下一个值，模拟生成一个环形链表，由于重复数的存在，必然会在重复数的地方开始绕环；
  - 使用快慢指针遍历数组，则他们最终会相遇在链表上的一个结点，在用一个指针从链表头开始，慢指针从当前位置开始，每次两个指针各进一步，当两个指针再次相遇时，即是环的入口，即重复数字的值；
  - 时间复杂度：$O(n)$， 空间复杂度：$O(1)$



### 378.有序矩阵中第k小的元素

>给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第k小的元素。
>请注意，它是排序后的第k小元素，而不是第k个元素。
>
>示例:
>
>matrix = [
>   [ 1,  5,  9],
>   [10, 11, 13],
>   [12, 13, 15]
>],
>k = 8,
>
>返回 13。
>说明:
>你可以假设 k 的值永远是有效的, 1 ≤ k ≤ n^2 。

- 解法：

```go
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix) - 1
	left, right := matrix[0][0], matrix[n][n] + 1
    for left < right { //log(right - left)次
		mid := left + ((right - left) >> 1)
        if nLessMids(&matrix, n, mid) >= k { // nlog(n)
			right = mid
		} else  {
			left = mid + 1
		}
	}
	return left
}

func nLessMids(matrix *[][]int, row, mid int) int {
	count := 0
	for i := 0; i <= row; i++ {
		if (*matrix)[i][row] <= mid {
			count += row + 1
		} else {
			for j:=0; j <= row && (*matrix)[i][j] <= mid; j++{
				count++
			}
		}
	}
	return count
}
```

- 思路：二分法：
  - 思路类似第668题；
  - 与第668题的区别：不能剪枝，矩阵只是在行和列上各自有序，并不是整体都有序，要依次判断每一行是否存在比 mid 小的值存入 count中；
  - 时间复杂度：$O(log(max - min) * nlog(n))$，空间复杂度：$O(1)$



### 448.找到所有数组中消失的数字

>给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
>
>找到所有在 [1, n] 范围之间没有出现在数组中的数字。
>
>您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
>
>示例:
>
>输入:
>[4,3,2,7,8,2,3,1]
>
>输出:
>[5,6]
>

- 解法一：

```go
func findDisappearedNumbers(nums []int) []int {
	size := len(nums)
	var res = make([]int, 0, size)
	if size == 0 {
		return nil
	}

	var temp = make([]bool, size)

	for _, v := range nums {
		temp[v - 1] = true
	}
	for k, v := range temp {
		if v != true {
			res = append(res, k + 1)
		}
	}
	return res
}
```

- 思路：
  - 引入一个辅助数组，遍历原数组，将数组中取到的元素值减一当成辅助数组的索引，并将该索引对应的值设为true，代表该数已经出现过；
  - 遍历辅助数组，找出值不为true的所有下标，加一后即可得到数组中没有出现的数；
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$
- 方法二：

```go
func findDisappearedNumbers(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return nil
	}

	var res = make([]int, 0, size)
	for _, value := range nums {
		idx := abs(value) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	for i, v := range nums {
		if v > 0 {
			res = append(res, i + 1)
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```

- 思路：
  - 直接在原数组上进行修改；
  - 要想利用原数组的空间存储数字是否出现的信息，同时又要保证能够取到原来的数值信息，考虑使用相反数；
  - 遍历数组，将取到的元素值取绝对值后减1，当做数组的下标，访问该下标，并将下标对应的值改为负数，遍历完成后，数组中，元素值依旧为正的值所对应的下标+1即是缺失的数；
  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$



### 668.乘法表中第k小的数

>几乎每一个人都用 乘法表。但是你能在乘法表中快速找到第k小的数字吗？
>
>给定高度m 、宽度n 的一张 m * n的乘法表，以及正整数k，你需要返回表中第k 小的数字。
>
>例 1：
>
>输入: m = 3, n = 3, k = 5
>输出: 3
>解释: 
>乘法表:
>1	2	3
>2	4	6
>3	6	9
>
>第5小的数字是 3 (1, 2, 2, 3, 3).
>例 2：
>
>输入: m = 2, n = 3, k = 6
>输出: 6
>解释: 
>乘法表:
>1	2	3
>2	4	6
>
>第6小的数字是 6 (1, 2, 2, 3, 4, 6).
>注意：
>
>m 和 n 的范围在 [1, 30000] 之间。
>k 的范围在 [1, m * n] 之间。
>

- 解法：

```go
func findKthNumber(m int, n int, k int) int {
	left, right := 1, m * n + 1 //注意，此处的left和right是实际的值，不是下标，right是右边界，左闭右开原则，所以要加1

	for left < right {
		mid := left + ((right - left) >> 1)
		if nLessMid(m, n, mid, k) >= k { //求出比mid值小的元素的个数，与k进行比较
			right = mid //不能写成 mid - 1，因为有可能count的值刚好就等于k
		} else {
			left = mid + 1 //这里的加1不能漏
		}
	}
	return left
}

func nLessMid(m, n, mid, k int) int {
	count := 0
	for i := 1; i <= minFunc(m, mid); i++ {
		count += minFunc(n, mid / i)
		if count >= k {
			return count
		}
	}
	return count
}

func minFunc(a, b int) int {
	if a < b {
		return a
	}

	return b
}
```

- 思路：二分法
  - left，right分别代表矩阵中的最小值和最大值，即矩阵的左上角和右下角的值；
  - 根据 left 和 right 的值求出 mid，计算矩阵中值小于 mid 的元素个数 count，比较 count 和目标值 k 就可知晓，要查找的元素到底落在了矩阵中的哪个区间内，修正 left 和 right；
  - 统计 count 的方法：第 i 行的最后一个数是 i * n，如果最后一个数小于等于中间数，则该行所有数都应该被统计进 count。即 $i * n \le mid$，两边除以 i，结果是一样的
  - 在统计 count 的时候存在两种可以剪枝的情况：
    - 如果行数 i 远大于mid，即 mid / i 的结果为0，则说明已经越过了mid，后续count相加都等于 count + 0，故不用遍历m行，只需遍历到 min(m, mid) 即可；
    - 如果 count 已经大于 k，则直接返回即可不用再往后遍历；
  - 时间复杂度：$O(mlog(m * n))$，空间复杂度：$O(1)$
- 同类题：378



### 786. 第K个最小的素数分数

>一个已排序好的表 A，其包含 1 和其他一些素数.  当列表中的每一个 p<q 时，我们可以构造一个分数 p/q 。
>
>那么第 k 个最小的分数是多少呢?  以整数数组的形式返回你的答案, 这里 answer[0] = p 且 answer[1] = q.
>
>示例:
>输入: A = [1, 2, 3, 5], K = 3
>输出: [2, 5]
>解释:
>已构造好的分数,排序后如下所示:
>1/5, 1/3, 2/5, 1/2, 3/5, 2/3.
>很明显第三个最小的分数是 2/5.
>
>输入: A = [1, 7], K = 1
>输出: [1, 7]
>注意:
>
>A 的取值范围在 2 — 2000.
>每个 A[i] 的值在 1 —30000.
>K 取值范围为 1 —A.length * (A.length - 1) / 2
>

- 解法：

```go
func kthSmallestPrimeFraction(A []int, K int) []int {
	var left float64= 0.0
	var right float64= 1.0
	size := len(A)

	for left < right {
		var mid = float64(left+right) / 2 //此处是float64，left + right不会溢出
		var max float64 = 0
		p, q, j, count := 0, 0, 1, 0

		for i := 0; i < size - 1; i++ { //注意i,j的取值，i比j要小1
			for j < size && float64(A[i]) > float64(A[j]) * mid {
				j++
			}
			if size == j {
				break
			}
			count += size - j

			temp := float64(A[i]) / float64(A[j])
			if temp > max {
				max = temp
				p = i
				q = j
			}
		}

		if count == K {
			return []int{A[p], A[q]}
		} else if count > K {
			right = mid
		} else {
			left = mid
		}
	}
	return []int{}
}
```

- 思路：

  - 构造虚拟数组，A[i]为分子，A[j]为分母

    >1/2 | 1/3 | 1/5 |
    >
    >NA | 2/3 | 2/5 |
    >
    >NA | NA | 3/5 |

  - 规律为：

    - 同一行从左往右值依次减小；
    - 同一列从下往上值依次减小；

  - 遍历每一行，找到第一个比 mid 小的元素，统计出每行比 mid 大的元素的个数，并将这个元素与 max 进行比较，更新 max 和坐标；

  - j 下标可以复用，上一行定位到的 j 对应到该行中第一个小于 mid 的元素，转到下一行时，由虚拟数组的构造规律可知，恰好是下一个位置的起点；

  - 当计数 count 