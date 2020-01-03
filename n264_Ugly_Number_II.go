package main

import (
	. "LeetCode-go/utils"
	"container/heap"
	"fmt"
	"math"
)

//方法一：动态规划
//Time：O(n)，Space：O(n)
func nthUglyNumber(n int) int {
	if n <= 0 {
		return -1
	}

	uglyNums := make([]int, n) //使用一个数组来保存生成的丑数
	uglyNums[0] = 1 //放入第一个丑数1
	p2, p3, p5 := 0, 0, 0 //使用三个游标，分别代表之后游标所指的数要乘以2,3,5
	for i := 1; i < n; i++ { //注意第一个丑数为1，已经生成，所以是从1开始
		//三个游标初始指向同一个数，分别乘以2,3,5后就得到了以该数为基准生成的三个丑数，从这三个数中取出最小的一个放入结果集中
		minNum := min(uglyNums[p2] * 2, uglyNums[p3] * 3, uglyNums[p5] * 5)
		uglyNums[i] = minNum
		//找出本轮生成丑数的游标，将该游标后移，复用未移动的游标生成的丑数
		if uglyNums[p2] * 2 == minNum {p2++}
		if uglyNums[p3] * 3 == minNum {p3++}
		if uglyNums[p5] * 5 == minNum {p5++}
	}
	return uglyNums[n - 1] //注意返回的是 n - 1
}

func min(a, b, c int) int {
	return Min(a, Min(b, c))
}


//方法二：最小堆
//由丑数的定义可知，丑数是只包含质因数2,3,5的数，则丑数的生成可以由1开始，乘以2,3,5，得到下一批丑数，再从这些数中取出最小
//的一个，乘以2,3,5，以此类推
//Time：O(n*log(n))，Space：O(n)
func nthUglyNumber1(n int) int {
	if n <= 0 { //边界条件：n <= 0时不满足丑数的定义
		return -1
	}

	temp := new(minheapUglyNum)
	heap.Init(temp)

	uglyNum := -1
	minVal := math.MaxInt64 //由于丑数增长的很快，所以为了防止超出范围，加上一个限制
	heap.Push(temp, 1)
	for n > 0 { //注意是小于号，没有等于
		for uglyNum == (*temp)[0] { //这一步是为了去重，如果堆顶元素等于上一轮迭代时使用的丑数，则直接舍去
			heap.Pop(temp)
		}
		uglyNum = heap.Pop(temp).(int) //取出堆顶元素作为本轮生成丑数的基数，依次乘以2,3,5，并将结果加入堆中
		if uglyNum * 2 <= minVal {
			heap.Push(temp, uglyNum * 2)
		}
		if uglyNum * 3 <= minVal {
			heap.Push(temp, uglyNum * 3)
		}
		if uglyNum * 5 <= minVal {
			heap.Push(temp, uglyNum * 5)
		}
		n--
	}
	return uglyNum
}

type minheapUglyNum []int

func (m minheapUglyNum) Len() int {
	return len(m)
}

func (m minheapUglyNum) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minheapUglyNum) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m *minheapUglyNum) Pop() interface{} {
	n := len(*m)
	res := (*m)[n - 1]
	*m = (*m)[: n - 1]
	return res
}

func (m *minheapUglyNum) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func main() {
	res := nthUglyNumber(10)
	fmt.Println(res)
}
